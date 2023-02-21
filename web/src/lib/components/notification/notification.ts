import { writable } from "svelte/store"
import { NotificationItem } from './notification-item'


var newItem = writable(new NotificationItem())
var counter = 0

export default {
    show: show,
    error: error,
    newItem: newItem,
}

function show(title: string, message: string, timeout?: number) {
    if (timeout == null) timeout = 2000
    newItem.set({ id: ++counter, title: title, message: message, timeout: timeout })
}

function error(message: string, timeout?: number) {
    show('Error', message, timeout)
}

