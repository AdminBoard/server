<script lang="ts">
	import { page } from '$app/stores';
	import api from '$lib/api';
	import DataInput from '$lib/components/data-input.svelte';
	import notification from '$lib/components/notification/notification';
	import { createEventDispatcher } from 'svelte';

	let dispatch = createEventDispatcher();

	let data: Record<string, any> = { id: 0, icon: '', name: '', status: '' };

	let updatingIcon = false,
		updatingName = false,
		updatingStatus = false,
		updatingPage = false;

	page.subscribe((page) => {
		let hash = page.url.hash;
		if (hash == '') return;

		let params = new URLSearchParams(hash.substring(1));
		const id = params.get('id');
		if (id == null) return;

		reload(+id);
	});

	function reload(id: number) {
		data.id = 0;
		data.icon = '';
		data.name = '';
		api.get('admin/menu?id=' + id).then((resp) => {
			if (resp.status == 0) {
				data = resp.data;
			} else notification.show('Error', resp.message);
		});
	}

	function callback(msg: string) {
		if (msg == null) {
			dispatch('change');
			return;
		}
		notification.show('Error', msg);
	}

	function updateStatus(status: string) {
		updatingStatus = true;
		api.postData('admin/menu/status?id=' + data.id, { status: status })
			.then((_) => {
				data.status = status;
				dispatch('change');
			})
			.catch((e) => notification.show('Error', e))
			.finally(() => (updatingStatus = false));
	}
</script>

{#if data.id > 0}
	<div class="details flex-column space-y-2">
		<div class="flex items-center space-x-2">
			<label for="">Status:</label>
			{data.status}
			{#if data.status == 'draft' || data.status == 'suspended'}
				<button
					on:click={() => updateStatus('active')}
					disabled={updatingStatus}>publish</button
				>
			{:else if data.status == 'active'}
				<button
					on:click={() => updateStatus('suspended')}
					disabled={updatingStatus}
				>
					suspend
				</button>
			{/if}
		</div>
		<div class="flex items-center space-x-2">
			<label for="icon">Icon:</label>
			<DataInput
				name="icon"
				bind:value={data.icon}
				saveUrl="/admin/menu/update?id={data.id}"
				saveCallback={callback}
				disabled={updatingIcon}
				class="flex-1"
			/>
		</div>
		<div class="flex items-center space-x-2">
			<label for="name">Name:</label>
			<DataInput
				name="name"
				bind:value={data.name}
				saveUrl="/admin/menu/update?id={data.id}"
				saveCallback={callback}
				disabled={updatingName}
				class="flex-1"
			/>
		</div>
		<div class="flex items-center space-x-2">
			<label for="name">Page ID:</label>
			<DataInput
				name="page"
				bind:value={data.page}
				saveUrl="/admin/menu/update?id={data.id}"
				saveCallback={callback}
				disabled={updatingPage}
				class="flex-1"
			/>
		</div>
		<div style="height: 600px;" />
		<!-- <div class="flex items-center space-x-2">
			<label for="page">Page:</label>
			<select name="page">
				<option value="1">Page 1</option>
				<option value="1">Page 2</option>
				<option value="1">Page 3</option>
				<option value="1">Page 4</option>
				<option value="1">Page 5</option>
				<option value="1">Page 6</option>
			</select>
			<DataInput
				name="name"
				bind:value={data.name}
				saveUrl="/admin/menu/update?id={data.id}"
				saveCallback={callback}
				disabled={updatingName}
			/>
		</div> -->
	</div>
{/if}

<style lang="scss">
	.details {
		& label {
			min-width: 100px;
		}
	}
</style>
