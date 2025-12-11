<template>
  <div class="analytics-dashboard">
    <header class="dashboard-header">
      <h1>Analytics Dashboard</h1>
      <div class="period-selector">
        <select v-model="selectedPeriod" @change="fetchStats">
          <option value="day">Last 24 hours</option>
          <option value="week">Last 7 days</option>
          <option value="month">Last 30 days</option>
          <option value="year">Last year</option>
        </select>
        <button @click="exportData" class="btn btn-secondary">
          <FeatherIcon name="download" size="16" />
          Export Data
        </button>
      </div>
    </header>

    <div class="stats-overview">
      <div class="stat-card">
        <div class="stat-header">
          <FeatherIcon name="eye" size="24" />
          <h3>Page Views</h3>
        </div>
        <div class="stat-value">{{ formatNumber(stats.total_page_views) }}</div>
        <div class="stat-change" :class="getChangeClass(pageViewsChange)">
          <FeatherIcon :name="pageViewsChange >= 0 ? 'trending-up' : 'trending-down'" size="16" />
          <span>{{ Math.abs(pageViewsChange) }}%</span>
        </div>
      </div>

      <div class="stat-card">
        <div class="stat-header">
          <FeatherIcon name="users" size="24" />
          <h3>Visitors</h3>
        </div>
        <div class="stat-value">{{ formatNumber(stats.total_visitors) }}</div>
        <div class="stat-change" :class="getChangeClass(visitorsChange)">
          <FeatherIcon :name="visitorsChange >= 0 ? 'trending-up' : 'trending-down'" size="16" />
          <span>{{ Math.abs(visitorsChange) }}%</span>
        </div>
      </div>

      <div class="stat-card">
        <div class="stat-header">
          <FeatherIcon name="activity" size="24" />
          <h3>Sessions</h3>
        </div>
        <div class="stat-value">{{ formatNumber(stats.total_sessions) }}</div>
        <div class="stat-change" :class="getChangeClass(sessionsChange)">
          <FeatherIcon :name="sessionsChange >= 0 ? 'trending-up' : 'trending-down'" size="16" />
          <span>{{ Math.abs(sessionsChange) }}%</span>
        </div>
      </div>

      <div class="stat-card">
        <div class="stat-header">
          <FeatherIcon name="clock" size="24" />
          <h3>Avg. Session</h3>
        </div>
        <div class="stat-value">{{ formatDuration(avgSessionDuration) }}</div>
        <div class="stat-change" :class="getChangeClass(sessionDurationChange)">
          <FeatherIcon :name="sessionDurationChange >= 0 ? 'trending-up' : 'trending-down'" size="16" />
          <span>{{ Math.abs(sessionDurationChange) }}%</span>
        </div>
      </div>
    </div>

    <div class="charts-grid">
      <div class="chart-card">
        <h3>Page Views Trend</h3>
        <div class="chart-container">
          <canvas ref="pageViewsChart"></canvas>
        </div>
      </div>

      <div class="chart-card">
        <h3>Visitor Sources</h3>
        <div class="chart-container">
          <canvas ref="sourcesChart"></canvas>
        </div>
      </div>
    </div>

    <div class="content-stats">
      <div class="popular-content">
        <h3>Most Popular Blog Posts</h3>
        <div v-if="stats.popular_blogs && stats.popular_blogs.length > 0" class="content-list">
          <div v-for="blog in stats.popular_blogs" :key="blog.content_id" class="content-item">
            <div class="content-info">
              <h4>{{ blog.title }}</h4>
              <p>{{ blog.views }} views</p>
            </div>
            <NuxtLink :to="`/blog/${blog.slug}`" class="btn btn-secondary">
              View
            </NuxtLink>
          </div>
        </div>
        <div v-else class="no-data">
          <FeatherIcon name="file-text" size="32" />
          <p>No blog views yet</p>
        </div>
      </div>

      <div class="popular-content">
        <h3>Most Played Games</h3>
        <div v-if="stats.popular_games && stats.popular_games.length > 0" class="content-list">
          <div v-for="game in stats.popular_games" :key="game.content_id" class="content-item">
            <div class="content-info">
              <h4>{{ game.title }}</h4>
              <p>{{ game.engagement }} plays</p>
            </div>
            <NuxtLink :to="`/games/${game.slug}`" class="btn btn-secondary">
              Play
            </NuxtLink>
          </div>
        </div>
        <div v-else class="no-data">
          <FeatherIcon name="play" size="32" />
          <p>No games played yet</p>
        </div>
      </div>
    </div>

    <div class="realtime-stats">
      <h3>Real-time Activity</h3>
      <div class="realtime-grid">
        <div class="realtime-card">
          <div class="realtime-value">{{ realtimeStats.active_users || 0 }}</div>
          <div class="realtime-label">Active Users</div>
        </div>
        <div class="realtime-card">
          <div class="realtime-value">{{ realtimeStats.page_views || 0 }}</div>
          <div class="realtime-label">Page Views (last hour)</div>
        </div>
        <div class="realtime-card">
          <div class="realtime-value">{{ formatDuration(avgPageTime) }}</div>
          <div class="realtime-label">Avg. Time on Page</div>
        </div>
      </div>

      <div v-if="realtimeStats.current_visits && realtimeStats.current_visits.length > 0" class="current-visits">
        <h4>Current Visits</h4>
        <div class="visits-list">
          <div v-for="visit in realtimeStats.current_visits.slice(0, 5)" :key="visit.time" class="visit-item">
            <div class="visit-page">
              <FeatherIcon name="globe" size="14" />
              <span>{{ truncatePath(visit.page) }}</span>
            </div>
            <div class="visit-time">{{ formatTimeAgo(visit.time) }}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useApi } from '~/composables/useApi'
import { Chart, registerables } from 'chart.js'

Chart.register(...registerables)

const { apiFetch } = useApi()

const selectedPeriod = ref('week')
const stats = ref<any>({})
const realtimeStats = ref<any>({})
const pageViewsChart = ref<HTMLCanvasElement | null>(null)
const sourcesChart = ref<HTMLCanvasElement | null>(null)

const pageViewsChange = ref(0)
const visitorsChange = ref(0)
const sessionsChange = ref(0)
const sessionDurationChange = ref(0)
const avgSessionDuration = ref(0)
const avgPageTime = ref(0)

let pageViewsChartInstance: Chart | null = null
let sourcesChartInstance: Chart | null = null

const fetchStats = async () => {
  try {
    const data = await apiFetch<any>(`/analytics/stats?period=${selectedPeriod.value}`)
    stats.value = data
    
    // Calculate changes (simplified - would compare with previous period)
    pageViewsChange.value = 12 // Example
    visitorsChange.value = 8
    sessionsChange.value = 5
    sessionDurationChange.value = -3
    avgSessionDuration.value = 180 // 3 minutes
    
    // Update charts
    updateCharts()
  } catch (error) {
    console.error('Failed to fetch analytics:', error)
  }
}

const fetchRealtimeStats = async () => {
  try {
    const data = await apiFetch<any>('/analytics/realtime')
    realtimeStats.value = data
    avgPageTime.value = 45 // Example: 45 seconds
  } catch (error) {
    console.error('Failed to fetch realtime stats:', error)
  }
}

const updateCharts = () => {
  if (!pageViewsChart.value || !sourcesChart.value) return
  
  // Destroy existing charts
  if (pageViewsChartInstance) {
    pageViewsChartInstance.destroy()
  }
  if (sourcesChartInstance) {
    sourcesChartInstance.destroy()
  }
  
  // Create page views chart
  pageViewsChartInstance = new Chart(pageViewsChart.value, {
    type: 'line',
    data: {
      labels: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
      datasets: [{
        label: 'Page Views',
        data: [120, 190, 300, 500, 200, 300, 450],
        borderColor: '#bd93f9',
        backgroundColor: 'rgba(189, 147, 249, 0.1)',
        tension: 0.4,
        fill: true
      }]
    },
    options: {
      responsive: true,
      plugins: {
        legend: {
          display: false
        }
      },
      scales: {
        y: {
          beginAtZero: true,
          grid: {
            color: 'rgba(255, 255, 255, 0.1)'
          },
          ticks: {
            color: 'rgba(255, 255, 255, 0.7)'
          }
        },
        x: {
          grid: {
            color: 'rgba(255, 255, 255, 0.1)'
          },
          ticks: {
            color: 'rgba(255, 255, 255, 0.7)'
          }
        }
      }
    }
  })
  
  // Create sources chart
  sourcesChartInstance = new Chart(sourcesChart.value, {
    type: 'doughnut',
    data: {
      labels: ['Direct', 'Referral', 'Search', 'Social'],
      datasets: [{
        data: [40, 30, 20, 10],
        backgroundColor: [
          '#bd93f9',
          '#8b8be9',
          '#50fa7b',
          '#ffb86c'
        ],
        borderColor: 'transparent'
      }]
    },
    options: {
      responsive: true,
      plugins: {
        legend: {
          position: 'bottom',
          labels: {
            color: 'rgba(255, 255, 255, 0.7)',
            padding: 20
          }
        }
      }
    }
  })
}

const exportData = async () => {
  try {
    const response = await fetch('/api/analytics/export?format=csv', {
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token')}`
      }
    })
    
    const blob = await response.blob()
    const url = window.URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = 'analytics_export.csv'
    document.body.appendChild(a)
    a.click()
    window.URL.revokeObjectURL(url)
    document.body.removeChild(a)
  } catch (error) {
    console.error('Failed to export data:', error)
  }
}

const formatNumber = (num: number) => {
  if (!num) return '0'
  if (num >= 1000000) return (num / 1000000).toFixed(1) + 'M'
  if (num >= 1000) return (num / 1000).toFixed(1) + 'K'
  return num.toString()
}

const formatDuration = (seconds: number) => {
  if (seconds < 60) return `${seconds}s`
  const minutes = Math.floor(seconds / 60)
  const remainingSeconds = seconds % 60
  return `${minutes}m ${remainingSeconds}s`
}

const formatTimeAgo = (time: string) => {
  const date = new Date(time)
  const now = new Date()
  const diffMs = now.getTime() - date.getTime()
  const diffMins = Math.floor(diffMs / 60000)
  
  if (diffMins < 1) return 'Just now'
  if (diffMins < 60) return `${diffMins}m ago`
  if (diffMins < 1440) return `${Math.floor(diffMins / 60)}h ago`
  return `${Math.floor(diffMins / 1440)}d ago`
}

const truncatePath = (path: string) => {
  if (path.length > 40) {
    return path.substring(0, 37) + '...'
  }
  return path
}

const getChangeClass = (change: number) => {
  return change >= 0 ? 'positive' : 'negative'
}

let refreshInterval: NodeJS.Timeout | null = null

onMounted(async () => {
  await Promise.all([
    fetchStats(),
    fetchRealtimeStats()
  ])
  
  // Refresh realtime stats every 30 seconds
  refreshInterval = setInterval(fetchRealtimeStats, 30000)
})

onUnmounted(() => {
  if (refreshInterval) {
    clearInterval(refreshInterval)
  }
  
  if (pageViewsChartInstance) {
    pageViewsChartInstance.destroy()
  }
  if (sourcesChartInstance) {
    sourcesChartInstance.destroy()
  }
})
</script>

<style scoped>
.analytics-dashboard {
  padding: 2rem;
  background: var(--theme-bg);
  min-height: 100vh;
}

.dashboard-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.dashboard-header h1 {
  color: var(--theme-primary);
  margin: 0;
}

.period-selector {
  display: flex;
  gap: 1rem;
  align-items: center;
}

.period-selector select {
  padding: 0.5rem 1rem;
  background: var(--theme-bg);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  font-family: inherit;
  cursor: pointer;
}

.stats-overview {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.stat-card {
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--theme-border);
  border-radius: 8px;
  padding: 1.5rem;
  transition: all 0.2s;
}

.stat-card:hover {
  border-color: var(--theme-primary);
  transform: translateY(-2px);
}

.stat-header {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 1rem;
}

.stat-header svg {
  color: var(--theme-primary);
}

.stat-header h3 {
  margin: 0;
  font-size: 1rem;
  color: var(--theme-fg);
  opacity: 0.8;
}

.stat-value {
  font-size: 2.5rem;
  font-weight: bold;
  color: var(--theme-primary);
  margin-bottom: 0.5rem;
}

.stat-change {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.9rem;
}

.stat-change.positive {
  color: var(--theme-success);
}

.stat-change.negative {
  color: var(--theme-error);
}

.charts-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.chart-card {
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--theme-border);
  border-radius: 8px;
  padding: 1.5rem;
}

.chart-card h3 {
  margin: 0 0 1rem 0;
  color: var(--theme-primary);
}

.chart-container {
  height: 300px;
  position: relative;
}

.content-stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.popular-content {
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--theme-border);
  border-radius: 8px;
  padding: 1.5rem;
}

.popular-content h3 {
  margin: 0 0 1rem 0;
  color: var(--theme-primary);
}

.content-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.content-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  background: rgba(0, 0, 0, 0.3);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  transition: all 0.2s;
}

.content-item:hover {
  border-color: var(--theme-primary);
}

.content-info h4 {
  margin: 0 0 0.25rem 0;
  color: var(--theme-fg);
}

.content-info p {
  margin: 0;
  font-size: 0.9rem;
  color: var(--theme-fg);
  opacity: 0.6;
}

.no-data {
  text-align: center;
  padding: 2rem;
  color: var(--theme-fg);
  opacity: 0.6;
}

.no-data svg {
  margin-bottom: 1rem;
}

.realtime-stats {
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--theme-border);
  border-radius: 8px;
  padding: 1.5rem;
}

.realtime-stats h3 {
  margin: 0 0 1rem 0;
  color: var(--theme-primary);
}

.realtime-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
  margin-bottom: 1.5rem;
}

.realtime-card {
  text-align: center;
  padding: 1.5rem;
  background: rgba(0, 0, 0, 0.3);
  border: 1px solid var(--theme-border);
  border-radius: 8px;
}

.realtime-value {
  font-size: 2rem;
  font-weight: bold;
  color: var(--theme-primary);
  margin-bottom: 0.5rem;
}

.realtime-label {
  font-size: 0.9rem;
  color: var(--theme-fg);
  opacity: 0.8;
}

.current-visits h4 {
  margin: 0 0 1rem 0;
  color: var(--theme-fg);
}

.visits-list {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.visit-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.75rem;
  background: rgba(0, 0, 0, 0.3);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  font-size: 0.9rem;
}

.visit-page {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: var(--theme-fg);
}

.visit-time {
  color: var(--theme-fg);
  opacity: 0.6;
}

@media (max-width: 768px) {
  .analytics-dashboard {
    padding: 1rem;
  }
  
  .dashboard-header {
    flex-direction: column;
    align-items: stretch;
    gap: 1rem;
  }
  
  .charts-grid {
    grid-template-columns: 1fr;
  }
  
  .chart-container {
    height: 250px;
  }
  
  .content-stats {
    grid-template-columns: 1fr;
  }
}
</style>