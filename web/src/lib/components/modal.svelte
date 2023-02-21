<script lang="ts" context="module">
	import api from '$lib/api';
	import { get } from '$lib/http';
	import { LayoutNode } from '$lib/layout/layout-node';
	import Layout from '$lib/layout/layout.svelte';
	import { writable, type Writable } from 'svelte/store';
	import notification from './notification/notification';

	let itemStore: Writable<any> = writable(null);
	let dataStore: Writable<any> = writable(null);

	export function show(node: LayoutNode) {
		itemStore.set(node);
	}

	export function popup(url: string, data?: any) {
		if (data == null) dataStore.set({});
		else dataStore.set(data);

		const loading = new LayoutNode();
		loading.tag = 'Loading';
		itemStore.set(loading);

		if (url != '') {
			api.get('page?path=' + url).then((resp) => {
				if (resp.status == 0) {
					itemStore.set(new LayoutNode().parse(resp.data.layout));
				} else {
					itemStore.set({});
					notification.error(resp.message);
				}
			});
		}
	}
</script>

<script lang="ts">
	let items: LayoutNode[] = [];
	let data: any = {};

	itemStore.subscribe((item) => {
		if (item != null && Object.keys(item).length > 0) {
			if (items.length > 0) {
				if (items[items.length - 1].tag == 'Loading') items.pop();
			}
			items = [...items, item];
		} else {
			items.pop();
			items = [...items];
		}
	});

	dataStore.subscribe((dt) => {
		data = dt;
	});

	function click(ev: any) {
		const detail = ev.detail;
		console.log(detail);

		switch (detail.action) {
			case 'submit':
				// console.log('submit');

				// 	const payload: Record<string, any> = {};
				// 	if (detail.actionFields != null) {
				// 		const split = detail.actionFields.split(',');
				// 		split.forEach((el: string) => {
				// 			el = el.trim();
				// 			payload[el.replaceAll('.', '_')] = getValue(data, el);
				// 		});
				// 	}
				// 	itemStore.set(new LayoutNode('Loading'));

				// 	api.post(detail.actionUrl, payload).then((resp) => {
				// 		itemStore.set({});
				// 		if (resp.status == 0) {
				// 			if (detail.onSuccess != null) {
				// 				notification.show(
				// 					'Sukses',
				// 					detail.onSuccess.message,
				// 					5000
				// 				);
				// 				switch (detail.onSuccess.action) {
				// 					case 'close':
				// 						itemStore.set({});
				// 						break;
				// 				}
				// 			}
				// 		} else {
				// 			notification.error(resp.message);
				// 		}
				// 	});
				break;

			case 'close':
				itemStore.set({});
				break;
			default:
				console.log(ev);
		}
	}

	function close(ev: any) {
		console.log(ev.detail);
		itemStore.set(null);
	}
</script>

{#if items.length > 0}
	{#each items as layout}
		<div class="modal">
			{#if layout.tag == 'Loading'}
				<div class="lds-ripple">
					<div />
					<div />
				</div>
			{:else}
				<div class="content">
					<Layout
						{data}
						{layout}
						on:click={click}
						on:close={close}
						on:process={() =>
							itemStore.set(new LayoutNode('Loading'))}
						on:finish={() => itemStore.set({})}
						on:error={(ev) => notification.error(ev.detail)}
					/>
				</div>
			{/if}
		</div>
	{/each}
{/if}

<style lang="scss">
	.modal {
		@apply absolute top-0 left-0 w-full h-full z-10;
		background-color: transparentize(#000, 0.5);

		& .content {
			@apply absolute bg-gray-100 space-y-1 w-fit h-fit drop-shadow-xl rounded overflow-clip;
			top: 50%;
			left: 50%;
			transform: translate(-50%, -50%);
			min-width: 400px;
			overflow: visible;
		}
	}

	.lds-ripple {
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		display: inline-block;
		position: relative;
		width: 80px;
		height: 80px;
	}
	.lds-ripple div {
		position: absolute;
		border: 4px solid #fff;
		opacity: 1;
		border-radius: 50%;
		animation: lds-ripple 1s cubic-bezier(0, 0.2, 0.8, 1) infinite;
	}
	.lds-ripple div:nth-child(2) {
		animation-delay: -0.5s;
	}
	@keyframes lds-ripple {
		0% {
			top: 36px;
			left: 36px;
			width: 0;
			height: 0;
			opacity: 0;
		}
		4.9% {
			top: 36px;
			left: 36px;
			width: 0;
			height: 0;
			opacity: 0;
		}
		5% {
			top: 36px;
			left: 36px;
			width: 0;
			height: 0;
			opacity: 1;
		}
		100% {
			top: 0px;
			left: 0px;
			width: 72px;
			height: 72px;
			opacity: 0;
		}
	}
</style>
