<script lang="ts">
    export { clazz as class };

    let clazz = '';

    export let width = '';
</script>

<div
    class="flex space-x-1 {clazz}"
    style={width == '' ? '' : 'width:' + width + 'rem'}
>
    <slot />
</div>
