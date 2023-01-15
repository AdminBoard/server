<script lang="ts">
	import api from '$lib/api';
	import Button from '$lib/components/button.svelte';
	import Content from '$lib/components/content.svelte';
	import { showPage } from '$lib/components/modal.svelte';
	import notification from '$lib/components/notification/notification';
	import Table from '$lib/components/table.svelte';
	import Titlebar from '$lib/components/titlebar.svelte';
	import { onMount } from 'svelte';

	let loading = false;
	let title = 'Pages';

	let columns: any[] = [
		{ name: 'name', label: 'ID' },
		{ name: 'url', label: 'URL' },
		{ name: 'title', label: 'Title' },
		{ name: 'description', label: 'Description' },
	];

	let rows: any[] = [];

	onMount(() => {
		loading = true;
		api.postData('admin/pages')
			.then((data) => {
				if (data != null) rows = data;
			})
			.catch((e) => notification.show('Error', e))
			.finally(() => (loading = false));
	});

	function click(ev: any) {
		showPage('', '/admin/pages/add');
	}
</script>

<Titlebar
	><span>{title}</span><Button class="material-icons" on:click={click}
		>add</Button
	></Titlebar
>

<Content>
	<Table {columns} {rows} />
</Content>
