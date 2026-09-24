<template>
  <el-form ref="formRef" :model="form" :rules="rules" label-position="top" @submit.prevent>
    <el-form-item label="技能名称" prop="title">
      <el-input v-model="form.title" placeholder="例如：毕业照人像摄影" maxlength="30" show-word-limit />
    </el-form-item>

    <div class="form-row">
      <el-form-item label="技能类别" prop="category">
        <el-select v-model="form.category" placeholder="请选择类别">
          <el-option v-for="item in SKILL_CATEGORIES" :key="item" :label="item" :value="item" />
        </el-select>
      </el-form-item>
      <el-form-item label="所在校区" prop="campus">
        <el-select v-model="form.campus" placeholder="请选择校区">
          <el-option v-for="item in CAMPUSES" :key="item" :label="item" :value="item" />
        </el-select>
      </el-form-item>
    </div>

    <el-form-item label="熟练度" prop="level">
      <el-slider v-model="form.level" :min="LEVEL_MIN" :max="LEVEL_MAX" show-input />
    </el-form-item>

    <el-form-item label="可交换时段" prop="timeSlots">
      <el-select v-model="form.timeSlots" multiple placeholder="选择你可以用来交换的时间段">
        <el-option v-for="slot in TIME_SLOTS" :key="slot" :label="slot" :value="slot" />
      </el-select>
    </el-form-item>

    <el-form-item label="想学技能" prop="wantedSkills">
      <el-select
        v-model="form.wantedSkills"
        multiple
        filterable
        allow-create
        default-first-option
        placeholder="选择或输入你想学的技能方向"
      >
        <el-option v-for="item in SKILL_CATEGORIES" :key="item" :label="item" :value="item" />
      </el-select>
    </el-form-item>

    <el-form-item label="期望回报（可多选）">
      <el-select v-model="form.rewards" multiple placeholder="选择期望获得的回报类型">
        <el-option v-for="item in REWARD_TYPES" :key="item" :label="item" :value="item" />
      </el-select>
    </el-form-item>

    <el-form-item label="技能描述">
      <el-input v-model="form.description" type="textarea" :rows="3" placeholder="简单介绍你能提供的帮助和内容" />
    </el-form-item>

    <el-form-item label="作品或证书（选填）">
      <el-input v-model="form.portfolio" placeholder="例如：12组校园人像作品" />
    </el-form-item>

    <el-alert v-if="errorMessage" :title="errorMessage" type="error" show-icon class="form-error" />

    <el-button type="primary" :loading="submitting" class="submit-button" @click="submit">
      发布到技能墙
    </el-button>
  </el-form>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import type { FormInstance, FormRules } from 'element-plus';
import {
  CAMPUSES,
  LEVEL_MAX,
  LEVEL_MIN,
  REWARD_TYPES,
  SKILL_CATEGORIES,
  TIME_SLOTS,
} from '../../constants/skill.constants';
import { AppException } from '../../errors/AppException';
import { logger } from '../../logger/logger';
import { publishSkill } from '../../services/skill.service';
import type { PublishSkillResult } from '../../types/domain';

const emit = defineEmits<{ published: [result: PublishSkillResult] }>();

const formRef = ref<FormInstance>();
const submitting = ref(false);
const errorMessage = ref('');

const initialForm = () => ({
  title: '',
  category: '',
  level: 60,
  campus: '',
  timeSlots: [] as string[],
  wantedSkills: [] as string[],
  rewards: [] as string[],
  description: '',
  portfolio: '',
});

const form = reactive(initialForm());

const rules: FormRules = {
  title: [{ required: true, message: '请填写技能名称', trigger: 'blur' }],
  category: [{ required: true, message: '请选择技能类别', trigger: 'change' }],
  campus: [{ required: true, message: '请选择所在校区', trigger: 'change' }],
  level: [{ required: true, message: '请设置熟练度', trigger: 'change' }],
  timeSlots: [{ required: true, type: 'array', min: 1, message: '请至少选择一个可交换时段', trigger: 'change' }],
  wantedSkills: [{ required: true, type: 'array', min: 1, message: '请至少填写一个想学的技能方向', trigger: 'change' }],
};

async function submit() {
  errorMessage.value = '';
  const valid = await formRef.value?.validate().catch(() => false);
  if (!valid) return;
  submitting.value = true;
  try {
    const result = await publishSkill({ ...form });
    emit('published', result);
    formRef.value?.resetFields();
    Object.assign(form, initialForm());
  } catch (err) {
    errorMessage.value = err instanceof AppException ? err.message : '发布失败，请稍后重试';
    logger.warn('publish skill failed', err);
  } finally {
    submitting.value = false;
  }
}
</script>

<style scoped>
.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
}
.form-error {
  margin-bottom: 14px;
}
.submit-button {
  width: 100%;
}
</style>
