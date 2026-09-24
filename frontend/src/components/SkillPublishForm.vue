<template>
  <el-form label-position="top" @submit.prevent>
    <div class="form-grid">
      <el-form-item label="技能名称" required>
        <el-input v-model="form.title" placeholder="例如：毕业照人像摄影" maxlength="30" show-word-limit />
      </el-form-item>
      <el-form-item label="技能类别" required>
        <el-select v-model="form.category" placeholder="选择技能类别">
          <el-option v-for="category in categories" :key="category" :label="category" :value="category" />
        </el-select>
      </el-form-item>
    </div>
    <el-form-item :label="`熟练度：${form.level}`" required>
      <el-slider v-model="form.level" :min="SKILL_LEVEL_MIN" :max="SKILL_LEVEL_MAX" :step="1" />
    </el-form-item>
    <el-form-item label="可交换时段" required>
      <el-select v-model="form.timeSlots" multiple collapse-tags placeholder="选择可交换的时间段">
        <el-option v-for="slot in EXCHANGE_TIME_SLOTS" :key="slot" :label="slot" :value="slot" />
      </el-select>
    </el-form-item>
    <el-form-item label="想学技能" required>
      <el-select
        v-model="form.wantedSkills"
        multiple
        filterable
        allow-create
        default-first-option
        collapse-tags
        placeholder="输入或选择想学的技能，如：吉他"
      >
        <el-option v-for="category in categories" :key="category" :label="category" :value="category" />
      </el-select>
    </el-form-item>
    <el-form-item label="补充描述（可选）">
      <el-input v-model="form.description" type="textarea" :rows="2" maxlength="120" placeholder="简单介绍你能提供的帮助" />
    </el-form-item>
    <el-button type="primary" :loading="submitting" @click="submit">发布技能</el-button>
  </el-form>

  <el-alert v-if="error" class="form-feedback" :title="error" type="error" show-icon :closable="false" />

  <template v-if="result">
    <el-alert v-if="result.waiting" class="form-feedback" :title="result.message" type="info" show-icon :closable="false" />
    <div v-else class="form-feedback">
      <el-alert :title="result.message" type="success" show-icon :closable="false" />
      <FeatureCard
        v-for="match in result.matches"
        :key="match.id"
        :title="`与 ${match.provider} 匹配成功`"
        :description="match.recommendation"
      >
        <template #tag><el-tag type="warning">匹配度 {{ match.score }}%</el-tag></template>
        <p class="muted">{{ match.offerSkill }} ↔ {{ match.wantedSkill }}</p>
        <div class="tag-row">
          <span class="muted">共同时段：</span>
          <el-tag v-for="slot in match.commonSlots" :key="slot" effect="plain">{{ slot }}</el-tag>
        </div>
      </FeatureCard>
    </div>
  </template>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import FeatureCard from './FeatureCard.vue';
import { publishSkill } from '../services/skill.service';
import { logger } from '../logger/logger';
import {
  EXCHANGE_TIME_SLOTS,
  PUBLISH_FAILED_FALLBACK,
  SKILL_LEVEL_DEFAULT,
  SKILL_LEVEL_MAX,
  SKILL_LEVEL_MIN,
} from '../constants/skill.constants';
import type { PublishSkillPayload, PublishSkillResult } from '../types/domain';

defineProps<{ categories: string[] }>();
const emit = defineEmits<{ published: [] }>();

const initialForm = (): PublishSkillPayload => ({
  title: '',
  category: '',
  level: SKILL_LEVEL_DEFAULT,
  timeSlots: [],
  wantedSkills: [],
  description: '',
});

const form = reactive<PublishSkillPayload>(initialForm());
const submitting = ref(false);
const error = ref('');
const result = ref<PublishSkillResult | null>(null);

async function submit() {
  error.value = '';
  result.value = null;
  submitting.value = true;
  try {
    result.value = await publishSkill({ ...form });
    Object.assign(form, initialForm());
    emit('published');
  } catch (err) {
    error.value = err instanceof Error ? err.message : PUBLISH_FAILED_FALLBACK;
    logger.warn('发布技能失败', err);
  } finally {
    submitting.value = false;
  }
}
</script>
