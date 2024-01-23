<script lang="ts">
  import {Paginator, type PaginationSettings} from '@skeletonlabs/skeleton';
  import {Table} from '@skeletonlabs/skeleton';

  const source = [['12222', 'asd', 'v', '2024/10/01']];

  let paginationSettings: PaginationSettings = {
    page: 0,
    limit: 15,
    size: source.length,
    amounts: [5, 10, 15],
  } satisfies PaginationSettings;

  let tableHeaders: string[] = ['', 'Name', 'Downloaded', 'Date'];

  $: sourceSliced = source.slice(paginationSettings.page * paginationSettings.limit, paginationSettings.page * paginationSettings.limit + paginationSettings.limit);
</script>

<div class="h-full w-full center flex-col gap-4">
  <Table source={{head: tableHeaders, body: sourceSliced}} class="w-3/4 h-2/4" />
  <Paginator bind:settings={paginationSettings} showFirstLastButtons={true} showPreviousNextButtons={true} showNumerals maxNumerals={1} class="focus:outline-none" />
</div>
