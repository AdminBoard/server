<script lang="ts">
	import { goto } from '$app/navigation';
	import api from '$lib/api';
	import type { SelectAction } from '$lib/class/select-action';
	import { createEventDispatcher, onMount } from 'svelte';
	import { showPage } from './modal.svelte';
	import Table from './table.svelte';

	export let source = '';
	export let columns: any[] = [];
	export let onSelect: SelectAction | null = null;
	export let selectable = false;

	let page = { offset: 0, count: 100 };

	let rows: any[] = [];
	const dispatch = createEventDispatcher();

	onMount(() => {
		if (source != '') {
			if (source.startsWith('/api/')) source = source.substring(4);
			api.post(source, { page: page }).then((resp) => {
				if (resp.status == 0) rows = resp.data;
			});
		}
	});

	function select(ev: any) {
		if (onSelect == null && !selectable) return;

		const detail = ev.detail;
		if (onSelect == null) {
			if (selectable) dispatch('select', ev.detail);
			return;
		}

		let url = onSelect.url == null ? '' : onSelect.url;

		switch (onSelect.action) {
			case 'popup':
				showPage('', url, detail);
				break;
			default:
				if (onSelect.url != '') {
					for (const key in detail) {
						url = url.replaceAll('%' + key + '%', detail[key]);
					}
					goto(url);
				}
		}
	}
</script>

<Table
	{columns}
	{rows}
	selectable={onSelect != null || selectable}
	on:select={select}
/>
<slot />
