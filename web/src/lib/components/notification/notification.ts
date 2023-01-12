import { writable } from "svelte/store"

export class NotificationItem {
    id: number = 0;
    title: string = '';
    message: string = '';
    timeout?: number = 2000;
}

var newItem = writable(new NotificationItem())
var counter = 0

export default {
    show: show,
    newItem: newItem,
}

function show(title: string, message: string, timeout?: number) {
    if (timeout == null) timeout = 2000;
    newItem.set({id: ++counter, title: title, message: message, timeout: timeout})
}

