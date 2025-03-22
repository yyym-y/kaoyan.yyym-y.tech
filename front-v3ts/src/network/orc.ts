import req from "./require";
import result from "@/model/utilModel/result";


export const latexOrc = (data:{"fileName" : string, "model" : number}) : Promise<result<any>> => {
    return req.post('/orc/latexOrc', data, {headers : {
        'Authorization' : localStorage.getItem("token")
    }}).then(res => {
        return res.data
    })
}