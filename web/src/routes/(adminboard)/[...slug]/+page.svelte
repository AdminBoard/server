<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import api from '$lib/api';
	import notification from '$lib/components/notification/notification';
	import { LayoutNode } from '$lib/layout/layout-node';
	import Layout from '$lib/layout/layout.svelte';

	import { onDestroy } from 'svelte';

	// let error = '';

	let layout: LayoutNode | null;

	let unsubscribe = page.subscribe((data) => {
		layout = null;

		api.get('page?path=/' + data.params.slug).then((resp) => {
			switch (resp.status) {
				case 0:
					layout = new LayoutNode().parse(resp.data.layout);
					break;
				case 404:
					goto('/404');
					break;
				default:
					notification.error(resp.message);
			}
		});
	});

	function error(ev: any) {
		notification.error(ev.detail);
	}

	onDestroy(unsubscribe);
</script>

<Layout {layout} on:error={error} />
