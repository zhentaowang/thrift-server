package com.wyun.thrift.server.business;

import com.alibaba.fastjson.JSONObject;

import java.util.HashMap;
import java.util.Map;

/**
 * User:lpf
 * Date:2017/6/10
 * Time:16:52
 */
public class BusinessServiceMap {
    public static Map<String, IBusinessService> serviceMap = new HashMap();

    public JSONObject handle(String serviceName, String operation, JSONObject request) {
        IBusinessService businessService=serviceMap.get(serviceName);
        JSONObject result=businessService.handle(operation,request);
        return result;
    }

    public void registerService(String serviceNmae,IBusinessService businessService){
        serviceMap.put(serviceNmae,businessService);
    }
}
