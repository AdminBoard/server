<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import api from '$lib/api';
	import { popup } from '$lib/components/modal.svelte';
	import { LayoutNode } from '$lib/layout/layout-node';
	import Layout from '$lib/layout/layout.svelte';
	import { onDestroy } from 'svelte';
	import AdminMenu from './menu/menu-page.svelte';

	let error = '';
	let kind = '';

	let component: any;
	let layout: LayoutNode | null;

	let unsubscribe = page.subscribe((data) => {
		kind = '';
		layout = null;
		component = null;
		api.get('page?path=/admin/' + data.params.slug).then((resp) => {
			if (resp.status == 0) {
				switch (resp.data.layout) {
					case 'admin.menu':
						component = AdminMenu;
						break;
					default:
						component = null;
						layout = new LayoutNode().parse(resp.data.layout);
				}
			} else {
				error = resp.message;
			}
		});
	});

	// function click(ev: any) {
	// 	const detail = ev.detail;

	// 	switch (detail.action) {
	// 		case 'popup':
	// 			if (detail.actionUrl != null) popup(detail.actionUrl);
	// 			break;
	// 		case 'open':
	// 			if (detail.actionUrl != null) goto(detail.actionUrl);
	// 			break;
	// 	}
	// }

	onDestroy(unsubscribe);
</script>

{#if error != ''}
	<h2>{error}</h2>
{/if}

{#if component == null}
	<Layout {layout} />
{:else}
	<svelte:component this={component} />
{/if}
