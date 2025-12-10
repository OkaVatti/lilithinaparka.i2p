<template>
  <section class="hero-section mb-12 text-center">
    <div class="mx-auto max-w-2xl">
      <!-- ASCII Branding header (static) -->
      <div class="mb-4 font-mono text-4xl md:text-6xl font-bold text-accent selection:bg-accent/30">
        <pre class="m-0 leading-tight">LilthinAparka</pre>
      </div>

      <p class="text-lg md:text-xl text-text-secondary mb-6 font-mono">
        &gt; personal blog-site + art gallery + portfolio + whatever i need it to be
      </p>

      <div
        class="mx-auto p-4 border rounded-lg bg-bg-secondary/50 ring-1 ring-accent/10"
        :class="['max-w-xl']"
      >
        <!-- Terminal-style table -->
        <div
          class="overflow-x-auto font-mono text-sm md:text-base"
          role="region"
          aria-label="profile table"
        >
          <table class="w-full table-auto text-left border-separate border-spacing-y-2">
            <thead>
              <tr>
                <th
                  class="w-1/3 pr-4 pb-2 text-xs uppercase text-text-secondary tracking-widest"
                >
                  Property (static)
                </th>
                <th class="pl-4 pb-2 text-xs uppercase text-text-secondary tracking-widest">
                  Value (dynamic)
                </th>
              </tr>
            </thead>

            <tbody>
              <tr
                v-for="prop in propsOrder"
                :key="prop.key"
                class="group"
              >
                <!-- static label -->
                <td class="pr-4 align-top">
                  <div class="px-2 py-1 rounded-sm bg-black/30 border border-accent/10">
                    <span class="text-accent/90 select-all">$&gt; {{ prop.label }}</span>
                  </div>
                </td>

                <!-- dynamic value fetched from API -->
                <td class="pl-4 align-top">
                  <div
                    class="px-3 py-2 rounded-sm bg-black/20 border border-accent/10 min-h-[2rem] flex items-center"
                  >
                    <template v-if="loading">
                      <span class="animate-pulse text-text-secondary">fetching…</span>
                    </template>

                    <template v-else-if="error">
                      <span class="text-red-400">error: {{ error }}</span>
                    </template>

                    <template v-else>
                      <span v-html="renderValue(prop.key)"></span>
                    </template>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>

          <!-- metadata footer -->
          <div class="mt-4 text-xs text-text-secondary font-mono">
            <span>Source:</span>
            <span class="ml-2 text-accent">{{ apiHost }}</span>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import useI2PFetch from '../../composables/useI2PFetch'

/**
 * Static ordering and labels for the table.
 * These are the "static properties" requested.
 */
const propsOrder = [
  { key: 'whoami', label: 'whoami' },
  { key: 'name', label: 'name' },
  { key: 'username', label: 'username' },
  { key: 'age', label: 'age' },
  { key: 'bio', label: 'bio' },
  { key: 'location', label: 'location' },
]

/**
 * Dynamic state
 */
const loading = ref(true)
const error = ref<string | null>(null)
const profile = ref<Record<string, unknown>>({})

/**
 * Read runtime-config I2P host and fetch the profile
 */
const { fetchFromI2P, apiHost } = useI2PFetch()

async function loadProfile() {
  loading.value = true
  error.value = null
  try {
    // endpoint path is configurable; change as needed
    const data = await fetchFromI2P('/api/profile')
    if (!data || typeof data !== 'object') {
      throw new Error('invalid payload')
    }
    profile.value = data as Record<string, unknown>
  } catch (err: any) {
    error.value = err?.message ?? String(err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  void loadProfile()
})

/**
 * Renderers for special formatting (keeps templates simple)
 */
function renderValue(key: string) {
  const v = profile.value[key]
  if (v == null) return `<span class="text-text-secondary">—</span>`

  // simple formatting:
  if (key === 'username') return `<span class="text-accent">@${escapeHtml(String(v))}</span>`
  if (key === 'bio') return `<em class="text-text-secondary">"${escapeHtml(String(v))}"</em>`
  if (key === 'age') return `<span class="text-accent">${escapeHtml(String(v))}</span>`

  return escapeHtml(String(v))
}

/**
 * Small helper to escape HTML in v-html usage
 */
function escapeHtml(unsafe: string) {
  return unsafe
    .replaceAll('&', '&amp;')
    .replaceAll('<', '&lt;')
    .replaceAll('>', '&gt;')
    .replaceAll('"', '&quot;')
    .replaceAll("'", '&#039;')
}
</script>

<style scoped>
/* Minimal accent tokens to match Tailwind palette variables you may define */
.text-accent { color: theme('colors.emerald.400', '#68D391'); }
.text-text-secondary { color: theme('colors.gray.400', '#9CA3AF'); }
.bg-bg-secondary { background-color: rgba(10, 10, 10, 0.45); }
</style>