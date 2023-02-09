<script lang="ts">
	import { goto } from '$app/navigation';
	import api from '$lib/api';
	import { popup } from '$lib/components/modal.svelte';
	import notification from '$lib/components/notification/notification';
	import { getValue } from '$lib/util';
	import { createEventDispatcher } from 'svelte';
	import type { LayoutNode } from './layout-node';
	import Renderer from './renderer.svelte';

	export let layout: LayoutNode | null = null;

	export let data: Record<string, any> = {};
	export let renderOnly = false;

	let dispatch = createEventDispatcher();
	let disabled = false;

	function action(act: any) {
		console.log(act);

		if (renderOnly) return;
		switch (act.action) {
			case 'open':
				goto(act.actionUrl);
				break;
			case 'popup':
				popup(act.actionUrl);
				break;
			case 'get':
				disabled = true;
				// api.get(act.actionUrl)
				// 	.then((resp) => {
				// 		if (resp.status == 0) {
				// 			if (act.onSuccess != null) action(act.onSuccess);
				// 		} else notification.show('Error', resp.message);
				// 	})
				// 	.catch((e) => notification.show('Error', e))
				// 	.finally(() => (disabled = false));
				break;
			case 'submit':
				if (act.actionFields != null && act.actionUrl != null) {
					disabled = true;
					const split = act.actionFields.split(',');
					const payload: Record<string, any> = {};
					split.forEach((el: string) => {
						el = el.trim();
						payload[el.replaceAll('.', '_')] = getValue(data, el);
					});

					let url = <string>act.actionUrl;

					let matches = [...url.matchAll(/\[([a-z0-9_]+)\]/gi)];
					matches.forEach((match) => {
						url = url.replaceAll(match[0], data[match[1]]);
					});
					api.post(url, payload)
						.then((resp) => {
							if (resp.status == 0) {
								disabled = false;
								if (act.onSuccess != null) {
									if (act.onSuccess.message != null)
										notification.show(
											'Success',
											act.onSuccess.message
										);
									dispatch('action', {
										action: act.onSuccess.action,
									});
								}
							} else notification.show('Error', resp.message);
						})
						.catch((e) => notification.show('Error', e));
				} else {
					notification.show('Error', 'Incomplete action defined');
				}

				break;

			default:
				dispatch('click', act);
		}
	}
</script>

<div class="content">
	{#if layout != null && layout.children != null}
		{#each layout.children as node}
			{#if node.tag != null}
				<Renderer {data} on:click={(ev) => action(ev.detail)} {node} />
			{:else}
				{@html node.text}
			{/if}
		{/each}
	{/if}
	<div class="overlay" class:invisible={!disabled} on:wheel|preventDefault>
		<div class="lds-ripple" class:invisible={!disabled}>
			<div />
			<div />
		</div>
	</div>
</div>

<style lang="scss">
	.content {
		@apply relative;
	}
	.overlay {
		@apply absolute top-0 w-full h-full;
		background-color: transparentize($color: #000000, $amount: 0.75);
	}

	.lds-ripple {
		@apply absolute top-5 right-10 inline-block w-10 h-10;
	}
	.lds-ripple div {
		@apply absolute opacity-100;
		border: 4px solid #fff;
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
