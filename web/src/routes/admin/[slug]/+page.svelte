<script lang="ts">
	import { page } from '$app/stores';
	import api from '$lib/api';
	import { popup, showPage } from '$lib/components/modal.svelte';
	import { LayoutNode } from '$lib/layout/layout-node';
	import Layout from '$lib/layout/layout.svelte';
	import { onDestroy } from 'svelte';
	import AdminMenu from './menu/menu-page.svelte';
	import AdminPages from './pages/pages-page.svelte';

	let error = '';
	let kind = '';

	let component: any;
	let layout: LayoutNode;

	function refresh(newPath: string) {
		error = '';
	}

	let unsubscribe = page.subscribe((data) => {
		kind = '';
		api.get('page?url=/admin/' + data.params.slug).then((resp) => {
			if (resp.status == 0) {
				switch (resp.data.layout) {
					case 'admin.menu':
						component = AdminMenu;
						break;
					case 'admin.pages':
						component = AdminPages;
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

	function click(ev: any) {
		const detail = ev.detail;

		// const params = detail.params;
		switch (detail.action) {
			case 'popup':
				if (detail.actionUrl != null) popup(detail.actionUrl);
				break;
		}
	}

	onDestroy(unsubscribe);
</script>

{#if error != ''}
	<h2>{error}</h2>
{/if}

{#if component == null}
	<Layout {layout} on:click={click} />
{:else}
	<svelte:component this={component} />
{/if}
