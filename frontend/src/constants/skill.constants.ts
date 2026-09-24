export const SKILL_CATEGORIES = ['摄影', '编程', '乐器', '外语', '平面设计', '健身指导'] as const;

export const TIME_SLOTS = [
  '周一晚',
  '周二晚',
  '周三晚',
  '周四晚',
  '周五晚',
  '周六上午',
  '周六下午',
  '周日上午',
  '周日下午',
  '周日全天',
] as const;

export const CAMPUSES = ['东校区', '西校区', '中心校区'] as const;

export const REWARD_TYPES = ['技能交换', '小额报酬', '请吃饭', '无偿'] as const;

export const LEVEL_MIN = 1;
export const LEVEL_MAX = 100;

export const SKILL_STATUS_MATCHED = '已匹配';
export const SKILL_STATUS_WAITING = '等待匹配';
