export default interface result<T = any> {
    code : number,
    msg : string,
    data : T
}