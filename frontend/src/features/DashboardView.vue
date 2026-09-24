<template>
  <main class="page-shell" v-loading="loading">
    <AppHeader :unread="overview?.metrics.unread ?? 0" />

    <section v-if="overview" class="metrics-grid">
      <MetricCard label="已发布技能" :value="overview.metrics.skills" />
      <MetricCard label="活跃需求" :value="overview.metrics.needs" />
      <MetricCard label="智能匹配" :value="overview.metrics.matches" />
      <MetricCard label="评价记录" :value="overview.metrics.reviews" />
    </section>

    <el-alert v-if="error" :title="error" type="error" show-icon />

    <section v-if="overview" class="workspace-grid">
      <div class="panel panel--full">
        <h2>发布我的技能</h2>
        <SkillPublishForm :categories="overview.categories" @published="loadOverview" />
      </div>

      <div class="panel">
        <h2>技能发布</h2>
        <FeatureCard v-for="skill in overview.skills" :key="skill.id" :title="skill.title" :description="skill.description">
          <template #tag>
            <el-tag>{{ skill.category }} {{ skill.level }}%</el-tag>
            <el-tag v-if="isWaitingMatch(skill)" type="info" effect="plain">{{ WAITING_MATCH_TAG }}</el-tag>
          </template>
          <div class="tag-row">
            <el-tag v-for="slot in skill.timeSlots" :key="slot" effect="plain">{{ slot }}</el-tag>
            <el-tag v-for="reward in skill.rewards" :key="reward" type="success" effect="plain">{{ reward }}</el-tag>
          </div>
          <div v-if="skill.wantedSkills?.length" class="tag-row">
            <el-tag v-for="wanted in skill.wantedSkills" :key="wanted" type="warning" effect="plain">想学 {{ wanted }}</el-tag>
          </div>
          <small>{{ skill.owner }} · {{ skill.campus }} · {{ skill.portfolio }}</small>
        </FeatureCard>
      </div>

      <div class="panel">
        <h2>需求浏览</h2>
        <el-table :data="overview.needs" size="small">
          <el-table-column prop="title" label="需求" min-width="170" />
          <el-table-column prop="category" label="类别" width="82" />
          <el-table-column prop="campus" label="校区" width="96" />
          <el-table-column prop="responses" label="响应" width="72" sortable />
        </el-table>
      </div>

      <div class="panel">
        <h2>智能匹配</h2>
        <FeatureCard v-for="match in overview.matches" :key="match.id" :title="`${match.provider} × ${match.learner}`" :description="match.recommendation">
          <template #tag><el-tag type="warning">{{ match.score }}%</el-tag></template>
          <p class="muted">{{ match.offerSkill }} ↔ {{ match.wantedSkill }}</p>
          <div class="tag-row">
            <el-tag v-for="slot in match.commonSlots" :key="slot">{{ slot }}</el-tag>
          </div>
        </FeatureCard>
      </div>

      <div class="panel">
        <h2>预约确认</h2>
        <el-timeline>
          <el-timeline-item v-for="item in overview.appointments" :key="item.id" :timestamp="item.time">
            <strong>{{ item.pair }}</strong>
            <p>{{ item.place }} · {{ item.status }}</p>
            <p class="muted">{{ item.agenda }}</p>
          </el-timeline-item>
        </el-timeline>
      </div>

      <div class="panel profile-panel">
        <div>
          <h2>个人主页与技能墙</h2>
          <h3>{{ overview.profile.name }}</h3>
          <p>{{ overview.profile.major }} · {{ overview.profile.creditLevel }}</p>
          <el-progress :percentage="overview.profile.creditScore" />
          <div class="tag-row">
            <el-tag v-for="skill in overview.profile.skillWall" :key="skill.id" effect="dark">{{ skill.title }} {{ skill.level }}%</el-tag>
          </div>
          <ul>
            <li v-for="item in overview.profile.history" :key="item">{{ item }}</li>
          </ul>
        </div>
        <RadarChart :radar="overview.profile.radar" />
      </div>

      <div class="panel">
        <h2>评价信用</h2>
        <FeatureCard v-for="review in overview.reviews" :key="review.id" :title="`${review.from} → ${review.to}`" :description="review.content">
          <template #tag><el-rate :model-value="review.rating" disabled size="small" /></template>
        </FeatureCard>
      </div>

      <div class="panel">
        <h2>消息通知</h2>
        <FeatureCard v-for="conversation in overview.messages" :key="conversation.id" :title="conversation.withUser" :description="conversation.messages.join(' / ')">
          <template #tag><el-badge :value="conversation.unread" /></template>
        </FeatureCard>
      </div>
    </section>
  </main>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import AppHeader from '../components/AppHeader.vue';
import FeatureCard from '../components/FeatureCard.vue';
import MetricCard from '../components/MetricCard.vue';
import RadarChart from '../components/RadarChart.vue';
import SkillPublishForm from '../components/SkillPublishForm.vue';
import { fetchOverview } from '../services/storage.service';
import { WAITING_MATCH_TAG } from '../constants/skill.constants';
import type { Overview, Skill } from '../types/domain';

const overview = ref<Overview | null>(null);
const loading = ref(true);
const error = ref('');

// 当前用户已被匹配覆盖的技能标题，用于判断哪些发布仍在等待
const matchedOwnTitles = computed(() => {
  const titles = new Set<string>();
  const me = overview.value?.profile.name;
  if (!me) return titles;
  for (const match of overview.value?.matches ?? []) {
    if (match.provider === me) titles.add(match.offerSkill);
    if (match.learner === me) titles.add(match.wantedSkill);
  }
  return titles;
});

function isWaitingMatch(skill: Skill): boolean {
  return skill.owner === overview.value?.profile.name && !matchedOwnTitles.value.has(skill.title);
}

async function loadOverview() {
  try {
    overview.value = await fetchOverview();
    error.value = '';
  } catch (err) {
    error.value = err instanceof Error ? err.message : '加载失败';
  } finally {
    loading.value = false;
  }
}

onMounted(loadOverview);
</script>
