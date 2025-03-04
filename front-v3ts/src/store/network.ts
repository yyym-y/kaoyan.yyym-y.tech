import { baseURL } from "@/network/require";
import { defineStore } from "pinia";

export const useServeURL = defineStore('url', {
	state: () => ({
		url : baseURL
	}),
	getters : {
		mergeURL: (state) => {
			return (path:string, subpath:string=""):string => {
				if (subpath != "" && !(subpath.endsWith("/") || subpath.endsWith("\\"))) {
					subpath += "/"
				}
				if(state.url.endsWith("/") || state.url.endsWith("\\"))
					return state.url + subpath + path
				else
					return state.url + "/" + subpath + path
			}
		}
	}
})