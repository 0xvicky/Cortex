export interface JobRecord {
  job_id: string
  repo_url: string
  created_at: string
}

export interface IngestRequest {
  repo_url: string
}

export interface IngestResponse {
  job_id: string
}

export interface JobsResponse {
  jobs: JobRecord[]
}

export interface RepoSummary {
  architecture?: string
  project_purpose?: string
  tech_stack?: string[]
  key_components?: Array<string | { name?: string; description?: string; [key: string]: any }>
  how_to_run?: string[] | string
  use_cases?: string[]
  [key: string]: any
}

export interface JobResult {
  job_id: string
  owner:string
  repo_name:string
  repo_url?: string
  status?: string
  repo_summary?: RepoSummary | null
}

export interface SummaryResponse {
  job: JobResult
}

export interface QueryRequest {
  job_id: string
  question: string
}

export interface QueryResponse {
  answer: string
}

export interface RepoOwner{
    owner:string
    repoName:string
}