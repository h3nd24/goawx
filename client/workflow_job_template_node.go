package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// WorkflowJobTemplateNodeService implements awx job template node apis.
type WorkflowJobTemplateNodeService struct {
	client *Client
}

// ListWorkflowJobTemplateNodesResponse represents `ListWorkflowJobTemplateNodes` endpoint response.
type ListWorkflowJobTemplateNodesResponse struct {
	Pagination
	Results []*WorkflowJobTemplateNode `json:"results"`
}

const workflowJobTemplateNodeAPIEndpoint = "/api/v2/workflow_job_template_nodes/"

// GetWorkflowJobTemplateNodeByID shows the details of a job template node.
func (jt *WorkflowJobTemplateNodeService) GetWorkflowJobTemplateNodeByID(id int, params map[string]string) (*WorkflowJobTemplateNode, error) {
	result := new(WorkflowJobTemplateNode)
	endpoint := fmt.Sprintf("%s%d/", workflowJobTemplateNodeAPIEndpoint, id)
	resp, err := jt.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// ListWorkflowJobTemplateNodes shows a list of job templates nodes.
func (jt *WorkflowJobTemplateNodeService) ListWorkflowJobTemplateNodes(params map[string]string) ([]*WorkflowJobTemplateNode, *ListWorkflowJobTemplateNodesResponse, error) {
	result := new(ListWorkflowJobTemplateNodesResponse)

	resp, err := jt.client.Requester.GetJSON(workflowJobTemplateNodeAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateWorkflowJobTemplateNode creates a job template node, without any pe exisiting nodes.
func (jt *WorkflowJobTemplateNodeService) CreateWorkflowJobTemplateNode(data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error) {
	result := new(WorkflowJobTemplateNode)
	mandatoryFields = []string{"workflow_job_template", "unified_job_template", "identifier"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := jt.client.Requester.PostJSON(workflowJobTemplateNodeAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateWorkflowJobTemplateNode updates a job template node.
func (jt *WorkflowJobTemplateNodeService) UpdateWorkflowJobTemplateNode(id int, data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error) {
	result := new(WorkflowJobTemplateNode)
	endpoint := fmt.Sprintf("%s%d", workflowJobTemplateNodeAPIEndpoint, id)
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

// DeleteWorkflowJobTemplateNode deletes a job template node.
func (jt *WorkflowJobTemplateNodeService) DeleteWorkflowJobTemplateNode(id int) (*WorkflowJobTemplateNode, error) {
	result := new(WorkflowJobTemplateNode)
	endpoint := fmt.Sprintf("%s%d", workflowJobTemplateNodeAPIEndpoint, id)

	resp, err := jt.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

func ValidateLinkType(linkType string) error {
	if linkType != "always_nodes" && linkType != "success_nodes" && linkType != "failure_nodes" {
		return fmt.Errorf("Currently the valid values for workflow job template node link type is [always_nodes, success_nodes, failure_nodes]")
	}
	return nil 
}

func (jt *WorkflowJobTemplateNodeService) AssociateWorkflowJobTemplateNodes(id int, linkType string, data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error) {
	err := ValidateLinkType(linkType)
	if err != nil {
		return nil, err
	}

	result := new(WorkflowJobTemplateNode)
	mandatoryFields = []string{"id"}
	validate, status := ValidateParams(data, mandatoryFields)
	data["associate"] = true
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	endpoint := fmt.Sprintf("%s%d/%s/", workflowJobTemplateNodeAPIEndpoint, id, linkType)
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

func (jt *WorkflowJobTemplateNodeService) DisAssociateWorkflowJobTemplateNodes(id int, linkType string, data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error) {
	err := ValidateLinkType(linkType)
	if err != nil {
		return nil, err
	}

	result := new(WorkflowJobTemplateNode)
	mandatoryFields = []string{"id"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	data["disassociate"] = true
	endpoint := fmt.Sprintf("%s%d/%s", workflowJobTemplateNodeAPIEndpoint, id, linkType)
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

func (jt *WorkflowJobTemplateNodeService) ListAssociatedWorkflowJobTemplateNodes(id int, linkType string, params map[string]string) ([]*WorkflowJobTemplateNode, *ListWorkflowJobTemplateNodesResponse, error) {
	result := new(ListWorkflowJobTemplateNodesResponse)
	err := ValidateLinkType(linkType)
	if err != nil {
		return nil, result, err
	}

	endpoint := fmt.Sprintf("%s%d/%s", workflowJobTemplateNodeAPIEndpoint, id, linkType)
	resp, err := jt.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

