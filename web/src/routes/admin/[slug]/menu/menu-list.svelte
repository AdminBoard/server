<script lang="ts">
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
  import api from '$lib/api';
  import DataInput from '$lib/components/data-input.svelte';
  import Menu from '$lib/components/menu.svelte';
  import notification from '$lib/components/notification/notification';
  import { onMount } from 'svelte';
  import { get } from 'svelte/store';
  import MenuAdd from './menu-add.svelte';

  let menu: any[] = [];
  let startAddMenu = false;
  let addSubmenu = new Map<number, string>();

  page.subscribe((_) => {
    selectMenu();
  });

  function selectMenu() {
    let hash = get(page).url.hash;
    if (hash == '') return;

    let params = new URLSearchParams(hash.substring(1));
    const id = params.get('id');
    if (id == null) return;

    menu.forEach((el, i) => {
      menu[i].selected = el.id == +id;
      if (el.children != null) {
        el.children.forEach((child: any, j: any) => {
          menu[i].children[j].selected = child.id == +id;
        });
      }
    });
  }

  function moveMenu(id1: number, id2: number) {
    let seq1 = menu[id1].sequence,
      seq2 = menu[id2].sequence;

    api
      .post('admin/menu/reorder', [
        { id: menu[id1].id, sequence: seq2 },
        { id: menu[id2].id, sequence: seq1 },
      ])
      .then((resp) => {
        if (resp.status == 0) {
          menu[id1].sequence = seq2;
          menu[id2].sequence = seq1;
          menu.sort((a, b) => {
            return a.sequence - b.sequence;
          });
        } else notification.show('Error', resp.message);
      })
      .catch((e) => notification.show('Error', e));
  }

  function saveMenu(
    parentID: number,
    ev: CustomEvent,
    callback?: (success: boolean) => void
  ) {
    let name = ev.detail;

    api
      .postData('admin/menu/add', { parent_id: parentID, name: name })
      .then((data) => {
        goto('/admin/menu#id=' + data.id);
        refresh();
        if (callback != null) callback(true);
      })
      .catch((e) => {
        notification.show('Error', e);
        if (callback != null) callback(false);
      })
      .finally(() => {
        if (parentID == 0) startAddMenu = false;
      });
  }

  export function refresh() {
    api.get('admin/menu').then((resp) => {
      if (resp.status == 0) {
        menu = resp.data;
        selectMenu();
      }
    });
  }

  onMount(() => refresh());
</script>

<div class="w-fit space-y-1">
  {#each menu as m, i}
    <!-- <div class="menu-group">
      <div class="w-full">
        <div class="flex relative w-full items-center">
          <Menu
            class="w-full"
            name={m.name}
            icon={m.icon}
            url="#id={m.id}"
            selected={m.selected}
          />
          {#if m.status == 'draft'}
            <i
              data-tooltip="draft"
              class="absolute tooltip-right cursor-default material-icons right-2 text-red-500 text-sm"
            >
              edit
            </i>
          {:else if m.status == 'suspended'}
            <i
              data-tooltip="suspended"
              class="absolute tooltip-right cursor-default material-icons right-2 text-yellow-500 text-sm"
            >
              pause
            </i>
          {/if}
        </div>
        {#each m.children as sub}
          <div class="flex relative w-full items-center">
            <Menu
              class="w-full"
              url="#id={sub.id}"
              name={sub.name}
              submenu={true}
              icon={sub.icon}
              selected={sub.selected}
            />
            {#if sub.status == 'draft'}
              <i
                data-tooltip="draft"
                class="absolute tooltip-right cursor-default material-icons right-2 text-red-500 text-sm"
              >
                edit
              </i>
            {:else if sub.status == 'suspended'}
              <i
                data-tooltip="suspended"
                class="absolute tooltip-right cursor-default material-icons right-2 text-yellow-500 text-sm"
              >
                pause
              </i>
            {/if}
          </div>
        {/each}
        <div class="relative add-submenu">
          <div class:invisible={!addSubmenu.has(m.id)}>
            <DataInput
              width="14"
              on:save={(ev) => {
                saveMenu(m.id, ev, () => addSubmenu.delete(m.id));
              }}
              on:cancel={() => {
                addSubmenu.delete(m.id);
                addSubmenu = addSubmenu;
              }}
              allowCancel={true}
            />

            <DataInput
              width="14"
              on:save={(ev) => {
                saveMenu(m.id, ev, () => addSubmenu.delete(m.id));
              }}
              on:cancel={() => {
                addSubmenu.delete(m.id);
                addSubmenu = addSubmenu;
              }}
              allowCancel={true}
            />
          </div>
          <button
            data-tooltip="Add submenu"
            class:hidden={addSubmenu.has(m.id)}
            class="btn-menu tooltip-right flex space-x-2 text-sm"
            on:click={() => {
              addSubmenu.set(m.id, '');
              addSubmenu = addSubmenu;
            }}
          >
            <i class="material-icons text-sm">add</i>
          </button>
        </div>
      </div>
      {#if m.sequence > 0}
        <div
          class="absolute right-2 top-2 tooltip-right"
          data-tooltip="Move up"
        >
          <button
            on:click={() => moveMenu(i - 1, i)}
            class="btn-sequence p-0 transition-colors rounded-full w-fit material-icons -rotate-90"
            >arrow_right</button
          >
        </div>
      {/if}
      {#if m.sequence != menu.length - 1}
        <div
          class="absolute right-2 bottom-2 tooltip-right"
          data-tooltip="Move down"
        >
          <button
            on:click={() => moveMenu(i, i + 1)}
            class="btn-sequence p-0 transition-colors rounded-full w-fit material-icons rotate-90"
            >arrow_right</button
          >
        </div>
      {/if}
    </div> -->
    <div class="menu-group">
      <Menu
        class="w-full"
        name={m.name}
        icon={m.icon}
        url="#id={m.id}"
        selected={m.selected}
      />
      {#if m.status == 'draft'}
        <i data-tooltip="draft" class="menu-badge material-icons ">edit</i>
      {:else if m.status == 'suspended'}
        <i data-tooltip="suspended" class="menu-badge material-icons ">pause</i>
      {/if}

      {#each m.children as sub}
        <div class="relative">
          <Menu
            class="w-full"
            url="#id={sub.id}"
            name={sub.name}
            submenu={true}
            icon={sub.icon}
            selected={sub.selected}
          />
          {#if sub.status == 'draft'}
            <i data-tooltip="draft" class="menu-badge material-icons">edit</i>
          {:else if sub.status == 'suspended'}
            <i data-tooltip="suspended" class="menu-badge material-icons "
              >pause</i
            >
          {/if}
        </div>
      {/each}

      <div class="ml-4">
        <MenuAdd
          parentID={m.id}
          on:success={() => refresh()}
          tooltip="Add submenu"
        />
      </div>

      {#if m.sequence > 0}
        <div class="move-up" data-tooltip="Move up">
          <button class="material-icons" on:click={() => moveMenu(i - 1, i)}
            >arrow_right</button
          >
        </div>
      {/if}
      {#if i != menu.length - 1}
        <div class="move-down" data-tooltip="Move down">
          <button class="material-icons" on:click={() => moveMenu(i, i + 1)}
            >arrow_right</button
          >
        </div>
      {/if}
    </div>
  {/each}

  <div class="menu-group">
    <MenuAdd parentID={0} on:success={() => refresh()} tooltip="Add menu" />
  </div>
</div>

<style lang="scss">
  @import '../../../../styles/colors';

  .menu-group {
    @apply relative border-2 rounded-lg flex-col p-1 items-center pl-8;
    min-width: 18em;
  }

  .menu-badge {
    @apply absolute tooltip-right cursor-default top-3 right-2 text-red-500 text-sm;
  }

  .btn-sequence {
    background-color: transparent;
    color: #666;
    box-shadow: 0 0 4px -2px #000;
    &:hover {
      box-shadow: 0 0 4px 0px $color-secondary;
      background-color: $color-secondary;
      color: $color-text;
    }
  }

  .move-up {
    @apply top-2;

    & button {
      @apply -rotate-90;
    }
  }

  .move-down {
    @apply bottom-2;

    & button {
      @apply rotate-90;
    }
  }

  .move-up,
  .move-down {
    @apply absolute w-fit left-2;
    & button {
      @apply p-0 transition-colors rounded-full;
      background-color: transparent;
      color: $color-primary;
    }
  }

  // .btn-menu {
  //   background-color: transparent;
  //   // text-align: left;
  //   position: absolute;
  //   top: 50%;
  //   transform: translateY(-50%);
  //   height: fit-content;

  //   &.hidden {
  //     display: none;
  //   }

  //   &:active {
  //     background-color: transparentize($color-secondary, 0.5);
  //     color: $color-text;
  //   }
  // }
  .add-submenu {
    margin: 0.5rem 0.5rem 0.5rem 2rem;
  }
</style>
