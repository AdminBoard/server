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
		updatingPath = false;

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
			} else notification.error(resp.message);
		});
	}

	function callback(msg: string) {
		if (msg == null) {
			dispatch('change');
			return;
		}
		notification.error(msg);
	}

	function updateStatus(status: string) {
		updatingStatus = true;
		api.postData('admin/menu/status?id=' + data.id, { status: status })
			.then((_) => {
				data.status = status;
				dispatch('change');
			})
			.catch((e) => notification.error(e))
			.finally(() => (updatingStatus = false));
	}
</script>

{#if data.id > 0}
	<div class="details">
		<div>
			<label for="status">Status:</label>
			{data.status}
			{#if data.status == 'draft' || data.status == 'suspended'}
				<button
					name="status"
					on:click={() => updateStatus('active')}
					disabled={updatingStatus}
				>
					publish
				</button>
			{:else if data.status == 'active'}
				<button
					name="status"
					on:click={() => updateStatus('suspended')}
					disabled={updatingStatus}
				>
					suspend
				</button>
			{/if}
		</div>
		<div>
			<label for="icon">Icon:</label>
			<DataInput
				name="icon"
				bind:value={data.icon}
				saveUrl="/admin/menu/update?id={data.id}"
				saveCallback={callback}
				disabled={updatingIcon}
			/>
		</div>
		<div>
			<label for="name">Name:</label>
			<DataInput
				name="name"
				bind:value={data.name}
				saveUrl="/admin/menu/update?id={data.id}"
				saveCallback={callback}
				disabled={updatingName}
			/>
		</div>
		<div>
			<label for="name">Page:</label>
			<DataInput
				name="page"
				bind:value={data.path}
				saveUrl="/admin/menu/update?id={data.id}"
				saveCallback={callback}
				disabled={updatingPath}
			/>
		</div>
	</div>
{/if}

<style lang="scss">
	.details {
		@apply flex-col space-y-2;
		& > div {
			@apply flex items-center space-x-2;
		}
		& label {
			min-width: 100px;
		}
	}
</style>
