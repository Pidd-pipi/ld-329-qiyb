<template>
  <div v-if="result" class="publish-result">
    <el-alert
      :title="result.message"
      :type="result.status === 'matched' ? 'success' : 'info'"
      show-icon
      :closable="false"
    />

    <template v-if="result.matches.length > 0">
      <FeatureCard
        v-for="match in result.matches"
        :key="match.id"
        :title="`${match.provider} × ${match.learner}`"
        :description="match.recommendation"
      >
        <template #tag><el-tag type="warning">匹配度 {{ match.score }}%</el-tag></template>
        <p class="muted">{{ match.offerSkill }} ↔ {{ match.wantedSkill }}</p>
        <div class="tag-row">
          <span class="muted">共同时段：</span>
          <el-tag v-for="slot in match.commonSlots" :key="slot">{{ slot }}</el-tag>
        </div>
      </FeatureCard>
    </template>

    <el-empty
      v-else
      description="暂时没有合适人选，技能已保留在技能墙上，仍在等待匹配"
      :image-size="90"
    />
  </div>

  <el-empty v-else description="填写左侧表单发布技能，发布后这里会立即显示匹配结果" :image-size="90" />
</template>

<script setup lang="ts">
import FeatureCard from '../../components/FeatureCard.vue';
import type { PublishSkillResult } from '../../types/domain';

defineProps<{ result: PublishSkillResult | null }>();
</script>

<style scoped>
.publish-result {
  display: grid;
  gap: 12px;
  align-content: start;
}
</style>
