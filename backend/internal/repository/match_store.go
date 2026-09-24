package repository

import "cyskillswap/internal/model"

// AddMatches 写入新生成的匹配并分配自增 ID，返回落库后的匹配
func AddMatches(matches []model.Match) []model.Match {
	if len(matches) == 0 {
		return matches
	}
	store.mu.Lock()
	defer store.mu.Unlock()
	saved := make([]model.Match, 0, len(matches))
	for _, match := range matches {
		match.ID = store.nextMatchID
		store.nextMatchID++
		store.matches = append(store.matches, match)
		saved = append(saved, match)
	}
	return saved
}
