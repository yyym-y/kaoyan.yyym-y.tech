import { createVedioReq } from "@/model/reqModel/createVedioReq";
import req from "./require";
import result from "@/model/utilModel/result";
import { vedioTableResp } from "@/model/respModel/vedioTableResp";


export const createVedio = (data:createVedioReq) : Promise<result<null>> => {
    return req.post('/admin/createVedio', data, {headers : {
        'Authorization' : localStorage.getItem("token")
    }}).then((res: { data: any; }) => {
        return res.data
    })
}

export const getDevelopContent = () : Promise<result<string>> => {
    return req.get('/admin/developLog', {headers : {
        'Authorization' : localStorage.getItem("token")
    }}).then((res: { data: any; }) => {
        return res.data
    })
}

export const saveDevelopContent = (data:{content : string}) : Promise<result<null>> => {
    return req.post('/admin/saveDevelopLog', data, {headers : {
        'Authorization' : localStorage.getItem("token")
    }}).then((res: { data: any; }) => {
        return res.data
    })
}

export const getVedioTable = () : Promise<result<vedioTableResp[]>> => {
    return req.get('/admin/getVedioTable', {headers : {
        'Authorization' : localStorage.getItem("token")
    }}).then((res: { data: any; }) => {
        return res.data
    })
}