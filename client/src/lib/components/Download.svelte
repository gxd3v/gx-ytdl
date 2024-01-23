<script lang="ts">
  import {currentTile} from '$lib/stores';
  import {focusTrap} from '@skeletonlabs/skeleton';
  import {get} from 'svelte/store';

  let isFocused: boolean = true;
</script>

<div class="h-full w-full center">
  <form class="w-3/4 card p-4 text-token space-y-4" use:focusTrap={isFocused}>
    <label class="label">
      <span>Youtube URL</span>
      <input class="focus:outline-none input pl-3" type="text" placeholder="https://www.youtube.com/watch?v=dQw4w9WgXcQ" />
    </label>

    <div class="center">
      <button class="btn variant-filled-surface">Download</button>
    </div>
  </form>
</div>
