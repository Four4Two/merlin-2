#!/bin/python3

# Usage: python3 prop12.py state_export_from_block_65838.json
# This script returns a csv of addresses and the amount of ufury to be paid to them 
# 65838 is the epoch detailed in Merlin prop 12
# state export should be obtained via version `v1.0.4`

import json
import functools
import sys

fury_issued = (300000000*1000000*0.45)/365

def dict_fury_rewards_by_gauge_id(data):
  # returns a map of the ufury rewards for each gauge id in poolincentives
  poolincentives = data["app_state"]["poolincentives"]
  gauge_weights = {x["gauge_id"] : int(x["weight"]) for x in poolincentives["distr_info"]["records"]}
  total_weight = int(poolincentives["distr_info"]["total_weight"])
  return {gid : (fury_issued*gauge_weights[gid]/total_weight) for gid in gauge_weights}

def dict_gauge_id_by_denom_duration(data):
  # returns a map from (denom, duration) to gauge ids
  incentives = data["app_state"]["incentives"]
  return {(x["distribute_to"]["denom"], x["distribute_to"]["duration"]) : x["id"] for x in incentives["gauges"]}

def dict_locked_by_gauge_id_by_address(data):
  # returns a map from gauge_id to (a map from address to amount locked in that gauge)
  accum = {}
  gauge_ids = dict_gauge_id_by_denom_duration(data)
  locks = data["app_state"]["lockup"]["locks"]

  for l in locks:
    amount = int(l["coins"][0]["amount"])
    denom = l["coins"][0]["denom"]
    duration = float(l["duration"][:-1])
    addr = l["owner"]

    if duration >= 86400:
      gid = gauge_ids[(denom, "86400s")]
      cur = accum.get(gid, {})
      cur[addr] = cur.get(addr, 0)+amount
      accum[gid] = cur

    if duration >= 604800:
      gid = gauge_ids[(denom, "604800s")]
      cur = accum.get(gid, {})
      cur[addr] = cur.get(addr, 0)+amount
      accum[gid] = cur
    
    if duration >= 1209600:
      gid = gauge_ids[(denom, "1209600s")]
      cur = accum.get(gid, {})
      cur[addr] = cur.get(addr, 0)+amount
      accum[gid] = cur
  return accum

def dict_rewards_by_address(data):
  fury_per_gauge = dict_fury_rewards_by_gauge_id(data)
  locked = dict_locked_by_gauge_id_by_address(data)
  gauge_totals = {gid : sum(locked[gid].values()) for gid in locked}

  #converts the {gauge_id -> {addr -> amount_locked}} dictionary into {gauge_id -> {addr -> usomo_issued}}
  # by multiplying by ufury issued to the gauge, and dividing by total locked for the gauge
  mint_by_gauge_addr = {gid : {addr : fury_per_gauge.get(gid, 0)*locked[gid][addr]/gauge_totals[gid] for addr in locked[gid]} for gid in locked}

  #merge dictionaries by adding when they share a key (assumes numerical values)
  #example: a={"foo":3, "bar":5} , b={"foo":1, "baz":8} => {"foo":4, "bar":5, "baz":8}
  merge_dict_by_sum = lambda a, b: {k : a[k] + b[k] if k in a and k in b else a.get(k, b.get(k)) for k in a.keys() | b.keys()}

  #convert {gauge_id -> {addr -> ufury_issued}} into {addr -> ufury_issued} by merging the dictionaries for each gauge id, adding the fury issued for each address
  return functools.reduce(merge_dict_by_sum, mint_by_gauge_addr.values())


def compute_sanity_check_and_export(data):
  rewards = dict_rewards_by_address(data)
  total_paid = sum(rewards.values())
  community_paid = dict_fury_rewards_by_gauge_id(data)["0"]

  print("expected total: ", int(fury_issued), "ufury")
  print("paid to LPers: ", int(total_paid), "ufury")
  print("paid to community pool: ", int(community_paid), "ufury")
  imbalance = int(fury_issued - (total_paid + community_paid))
  print("imbalance: ", imbalance)
  assert(imbalance == 0)

  f = open("prop_12_payments.csv","w")
  f.write("\n".join([addr+", "+str(int(rewards[addr])) for addr in sorted(rewards)]))
  f.close()



if __name__ == "__main__":
  fname = sys.argv[1]
  f = open(fname)
  data = json.loads(f.read())
  compute_sanity_check_and_export(data)