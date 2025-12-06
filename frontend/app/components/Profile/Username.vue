<script setup lang="ts">
import type { Profile } from '../../../types';

defineProps<{
  profile: Profile
}>()

</script>

<template>
    <div>
        <span>{{ profile.username }}</span>
    </div>
</template>