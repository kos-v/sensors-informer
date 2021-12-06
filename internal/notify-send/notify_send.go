package notify_send

type Urgency string

const (
	UrgencyLow      Urgency = "low"
	UrgencyNormal   Urgency = "normal"
	UrgencyCritical Urgency = "critical"
)

func IsValidUrgency(v Urgency) bool {
	return v == UrgencyLow ||
		v == UrgencyNormal ||
		v == UrgencyCritical
}
