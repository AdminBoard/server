<script lang="ts">
	import { page } from '$app/stores';
	import api from '$lib/api';
	import { showPage } from '$lib/components/modal.svelte';
	import { LayoutNode } from '$lib/layout/layout-node';
	import Layout from '$lib/layout/layout.svelte';

	import { onMount } from 'svelte';
	import { get } from 'svelte/store';

	let error = '';

	let layout: LayoutNode;

	onMount(() => {
		api.get('page?url=/' + get(page).params.slug).then((resp) => {
			if (resp.status == 0) {
				layout = new LayoutNode().parse(resp.data.layout);
			} else {
				error = resp.message;
			}
		});
	});

	function click(ev: any) {
		const detail = ev.detail;
		const params = detail.params;
		switch (detail.action) {
			case 'popup':
				showPage(params.title, params.url);
				break;
		}
	}
</script>

<Layout {layout} on:click={click} />
