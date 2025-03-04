import { notePath } from "@/model/bussModel/notePath";
import req from "./require";
import result from "@/model/utilModel/result";

export const getNotePath = ():Promise<result<notePath[]>> => {
    return req.get('/note/notePath').then(res => {
        return res.data
    })
}

export const getNoteContent = (params:{"notePath" : string}):Promise<result<string>> => {
    return req.get('/note/noteContent', {params}).then(res => {
        return res.data
    })
}

export const saveContent = (data:{ notePath:string, content:string }) : Promise<result<null>> => {
    return req.post('/note-admin/saveContent', data, {headers : {
        'Authorization' : localStorage.getItem("token")
    }}).then((res: { data: any; }) => {
        return res.data
    })
}

export const fileChange = (data:{ action:number, origin:string, final:string }) : Promise<result<null>> => {
    return req.post('/note-admin/fileChange', data, {headers : {
        'Authorization' : localStorage.getItem("token")
    }}).then((res: { data: any; }) => {
        return res.data
    })
}