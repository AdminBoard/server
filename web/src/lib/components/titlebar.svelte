<div
	class="titlebar flex items-center bg-primary-400 shadow-xs text-lg p-2 min-h-fit space-x-2"
>
	<slot />
</div>

<style lang="scss">
	.titlebar :global(.material-icons) {
		width: 32px;
		padding: 4px 2px;
	}
</style>
