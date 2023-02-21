<script lang="ts">
	import { page } from '$app/stores';
	import api from '$lib/api';
	import { LayoutNode } from '$lib/layout/layout-node';
	import Layout from '$lib/layout/layout.svelte';
	import { onDestroy } from 'svelte';

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
						// component = AdminMenu;
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
