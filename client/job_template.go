package awx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// JobTemplateService implements awx job template apis.
type JobTemplateService struct {
	client *Client
}

// ListJobTemplatesResponse represents `ListJobTemplates` endpoint response.
type ListJobTemplatesResponse struct {
	Pagination
	Results []*JobTemplate `json:"results"`
}

const jobTemplateAPIEndpoint = "/api/v2/job_templates/"

// GetJobTemplateByID shows the details of a job template.
func (jt *JobTemplateService) GetJobTemplateByID(id int, params map[string]string) (*JobTemplate, error) {
	result := new(JobTemplate)
	endpoint := fmt.Sprintf("%s%d/", jobTemplateAPIEndpoint, id)
	resp, err := jt.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// ListJobTemplates shows a list of job templates.
func (jt *JobTemplateService) ListJobTemplates(params map[string]string) ([]*JobTemplate, *ListJobTemplatesResponse, error) {
	result := new(ListJobTemplatesResponse)
	resp, err := jt.client.Requester.GetJSON(jobTemplateAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// Launch lauchs a job with the job template.
func (jt *JobTemplateService) Launch(id int, data map[string]interface{}, params map[string]string) (*JobLaunch, error) {
	result := new(JobLaunch)
	endpoint := fmt.Sprintf("%s%d/launch/", jobTemplateAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	// in case invalid job id return
	if result.Job == 0 {
		return nil, errors.New("invalid job id 0")
	}

	return result, nil
}

// CreateJobTemplate creates a job template
func (jt *JobTemplateService) CreateJobTemplate(data map[string]interface{}, params map[string]string) (*JobTemplate, error) {
	result := new(JobTemplate)
	mandatoryFields = []string{"name", "job_type", "inventory", "project"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PostJSON(jobTemplateAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateJobTemplate updates a job template
func (jt *JobTemplateService) UpdateJobTemplate(id int, data map[string]interface{}, params map[string]string) (*JobTemplate, error) {
	result := new(JobTemplate)
	endpoint := fmt.Sprintf("%s%d", jobTemplateAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteJobTemplate deletes a job template
func (jt *JobTemplateService) DeleteJobTemplate(id int) (*JobTemplate, error) {
	result := new(JobTemplate)
	endpoint := fmt.Sprintf("%s%d", jobTemplateAPIEndpoint, id)

	resp, err := jt.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DisAssociateCredentials remove Credentials form an awx job template
func (jt *JobTemplateService) DisAssociateCredentials(id int, data map[string]interface{}, params map[string]string) (*JobTemplate, error) {
	result := new(JobTemplate)
	endpoint := fmt.Sprintf("%s%d/credentials/", jobTemplateAPIEndpoint, id)
	data["disassociate"] = true
	mandatoryFields = []string{"id"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := jt.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// AssociateCredentials  adding credentials to JobTemplate.
func (jt *JobTemplateService) AssociateCredentials(id int, data map[string]interface{}, params map[string]string) (*JobTemplate, error) {
	result := new(JobTemplate)

	endpoint := fmt.Sprintf("%s%d/credentials/", jobTemplateAPIEndpoint, id)
	data["associate"] = true
	mandatoryFields = []string{"id"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := jt.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

func (jt *JobTemplateService) ListCredentials(id int, params map[string]string) ([]*Credential,
	*ListCredentialsResponse,
	error) {
	result := new(ListCredentialsResponse)
	endpoint := fmt.Sprintf("%s%d/credentials/", jobTemplateAPIEndpoint, id)
	resp, err := jt.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// Survey is attached to a job template, the id here is the job_template ID
func (jt *JobTemplateService) PostSurvey(id int, data Survey, params map[string]string) (*Survey, error) {
	result := new(Survey)
	endpoint := fmt.Sprintf("%s%d/survey_spec/", jobTemplateAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// the survey is attached to a job template, so we query the entire survey of a job template by its ID
func (jt *JobTemplateService) GetSurveyByJobTemplateId(id int, params map[string]string) (*Survey, error) {
	result := new(Survey)
	endpoint := fmt.Sprintf("%s%d/survey_spec/", jobTemplateAPIEndpoint, id)
	resp, err := jt.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// Delete the whole survey. AWX doesn't really have delete individual survey question, but re-POST everytime there is a change
func (jt *JobTemplateService) DeleteSurvey(id int) (*Survey, error) {
	result := new(Survey)
	endpoint := fmt.Sprintf("%s%d/survey_spec", jobTemplateAPIEndpoint, id)

	resp, err := jt.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
