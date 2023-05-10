import { goto } from "$app/navigation"
import { writable } from "svelte/store"
import api from "./api"
import md5 from "./md5"

let menu = writable([])
let username = ''
export let loggedIn = writable(false)

export default {
	username: username,
	load: load,
	login: login,
	logout: logout,
	changePassword: changePassword,
	menu: menu
}

function load(): Promise<any> {
	return new Promise<any>((resolve, reject) => {
		api.get('/user/session').then(resp => {
			if (resp.status < 100) {
				username = resp.data?.username
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

function login(u: string, p: string): Promise<boolean> {
	const time = Math.floor(new Date().getTime() / 1000)
	u = u.trim().toLowerCase()
	username = u

	const secret = md5(u + p)

	return new Promise<boolean>((resolve, reject) => {
		api.post('/user/login', { username: u, signature: md5(secret + time), time: time })
			.then(resp => {
				if (resp.status == 0) {
					loggedIn.set(true)
					refreshMenu()
					resolve(true)
				} else if (resp.status == 1) resolve(false)
				else reject(resp.message)
			}).catch(e => {
				reject(e)
			})
	})
}

function logout(): Promise<null> {
	return new Promise<null>(_ => {
		api.get('/user/logout').finally(() => {
			loggedIn.set(false)
			menu.set([])
			goto('/user/login')
		})
	})
}

function changePassword(password: string, confirm: string): Promise<boolean> {
	return new Promise<boolean>((resolve, reject) => {
		api.post('/user/change_password', { password: password, confirm: confirm }).then(resp => {
			if (resp.status == 0) {
				loggedIn.set(true)
				refreshMenu()
				resolve(true)
			}
			else reject(resp.message)
		})
	})
}

function refreshMenu() {
	api.get('menu').then((resp) => {
		if (resp.status == 0) menu.set(resp.data)
		else console.log('Error refreshMenu:' + resp.message)
	})
}