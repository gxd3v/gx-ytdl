<script lang="ts">
  import {message, socket} from '$lib/stores';
  import {CreateSessionResponse} from '$lib/protos/messages';

  let id: string = '-';

  message.subscribe((message: string) => {
    try {
      const data: CreateSessionResponse = CreateSessionResponse.fromJsonString(message);
      id = data.sessionId;
      localStorage.setItem('session', id);
    } catch (e) {
    }
  });

</script>

<div class="text-center label text-xs p-3">
  <span>Your session ID: {id}</span>
</div>
