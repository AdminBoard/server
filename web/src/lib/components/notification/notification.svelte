<script lang="ts">
	import notification, { NotificationItem } from './notification';

	let items: NotificationItem[] = [];

	notification.newItem.subscribe((item) => {
		if (item.title != null && item.title != '') {
			items = [...items, item];
			setTimeout(() => {
				items = items.filter((i) => {
					return i.id != item.id;
				});
			}, item.timeout);
		}
	});
</script>

{#if items.length > 0}
	<div class="fixed top-0 z-10 right-0 p-2 space-y-2 max-w-md">
		{#each items as item}
			<div class="item rounded shadow-lg bg-gray-700 border-l-4 border-blue-500 text-white p-4">
				<p class="font-bold">{item.title}</p>
				<p>{item.message}</p>
			</div>
		{/each}
	</div>
{/if}

<style lang="scss">
	.item {
		min-width: 300px;
	}
</style>
