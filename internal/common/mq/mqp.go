package mq

type JobLinkMessagePayload struct {
	Location   string `json:"location"`
	LocationId string `json:"location_id"`
	JobLink    string `json:"job_link"`
}

func CreateJobLinkMessagePayload(location string, locationId string, jobLink string) JobLinkMessagePayload {
	return JobLinkMessagePayload{Location: location, LocationId: locationId, JobLink: jobLink}
}
