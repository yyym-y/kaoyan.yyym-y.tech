import { vedio } from "@/model/bussModel/vedio";
import req from "./require";
import result from '@/model/utilModel/result'
import { playInfo } from "@/model/bussModel/playInfo";

export const getVedioById = (params:{"vedioId" : number}) : Promise<result<vedio>> => {
    return req.get('/vedio/getVedioById', {params, headers : {
        'Authorization' : localStorage.getItem("token")
    }}).then(res => {
        return res.data
    })
}

export const getPlayList = (data:{"vedioId" : number}) : Promise<result<playInfo[]>> => {
    return req.post('/vedio/getPlayList', data, {headers : {
        'Authorization' : localStorage.getItem("token")
    }}).then(res => {
        return res.data
    })
}