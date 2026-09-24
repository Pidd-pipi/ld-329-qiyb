import { AppException } from '../errors/AppException';
import type { PublishSkillPayload, PublishSkillResult } from '../types/domain';

const API_BASE = '/api';

interface ErrorBody {
  code?: string;
  message?: string;
}

export async function publishSkill(payload: PublishSkillPayload): Promise<PublishSkillResult> {
  const response = await fetch(`${API_BASE}/skills`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  });
  const body = (await response.json()) as PublishSkillResult & ErrorBody;
  if (!response.ok) {
    throw new AppException(body.message ?? '发布失败，请稍后重试', body.code);
  }
  return body;
}
