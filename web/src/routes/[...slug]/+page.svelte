<script lang="ts">
	import { page } from '$app/stores';
	import api from '$lib/api';
	import { LayoutNode } from '$lib/layout/layout-node';
	import Layout from '$lib/layout/layout.svelte';

	import { onDestroy } from 'svelte';

	let error = '';

	let layout: LayoutNode | null;

	let unsubscribe = page.subscribe((data) => {
		layout = null;

		api.get('page?path=/' + data.params.slug).then((resp) => {
			if (resp.status == 0) {
				layout = new LayoutNode().parse(resp.data.layout);
			} else {
				error = resp.message;
			}
		});
	});

	onDestroy(unsubscribe);
</script>

<Layout {layout} />
