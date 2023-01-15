import { goto } from "$app/navigation"
import { writable } from "svelte/store"
import api from "./api"
import md5 from "./md5"

let menu = writable([])
export let loggedIn = writable(false)

export default {
	load: load,
	login: login,
	logout: logout,
	menu: menu,
	loggedIn: loggedIn
}

export function load(): Promise<any> {
	return new Promise<any>((resolve, reject) => {
		api.get('session').then(resp => {
			if (resp.status < 100) {
				if (resp.status == 0) loggedIn.set(true)
				refreshMenu()
				resolve(resp)
			} else {
				loggedIn.set(false)
				menu.set([])
				reject(resp.message)
			}
		}).catch(e => {
			reject(e)
		})
	})
}

export function login(username: string, password: string): Promise<boolean> {
	const time = Math.floor(new Date().getTime() / 1000)

	const secret = md5(username + password)

	return new Promise<boolean>((resolve, reject) => {
		api.post('login', { username: username, signature: md5(secret + time), time: time })
			.then(resp => {
				if (resp.status == 0) {
					loggedIn.set(true)
					refreshMenu()
					resolve(true)
				}
				else reject(resp.message)
			}).catch(e => {
				reject(e)
			})
	})
}

function logout(): Promise<null> {
	return new Promise<null>(_ => {
		api.get('logout').finally(() => {
			loggedIn.set(false)
			menu.set([])
			goto('/user/login')
		})
	})
}

function refreshMenu() {
	api.get('menu').then((resp) => {
		if (resp.status == 0) menu.set(resp.data)
		else console.log('Error refreshMenu:' + resp.message)
	})
}