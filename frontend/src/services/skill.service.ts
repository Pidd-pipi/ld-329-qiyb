import { AppException } from '../errors/AppException';
import { PUBLISH_FAILED_FALLBACK } from '../constants/skill.constants';
import { logger } from '../logger/logger';
import type { PublishSkillPayload, PublishSkillResult } from '../types/domain';

const API_BASE = '/api';

interface ApiErrorBody {
  code?: string;
  message?: string;
}

export async function publishSkill(payload: PublishSkillPayload): Promise<PublishSkillResult> {
  const response = await fetch(`${API_BASE}/skills`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  });
  if (!response.ok) {
    let code = 'PUBLISH_FAILED';
    let message = PUBLISH_FAILED_FALLBACK;
    try {
      const body = (await response.json()) as ApiErrorBody;
      if (body.message) message = body.message;
      if (body.code) code = body.code;
    } catch {
      logger.warn('publish skill: 无法解析失败响应');
    }
    throw new AppException(message, code);
  }
  return response.json() as Promise<PublishSkillResult>;
}
