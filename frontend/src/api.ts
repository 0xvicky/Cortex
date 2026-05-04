import type { RepoSummary } from './types'

const BASE_URL = import.meta.env.VITE_API_BASE_URL;
console.log(BASE_URL)
async function apiRequest<T>(path: string, options: RequestInit = {}) {
  const response = await fetch(`${BASE_URL}${path}`, {
    headers: {
      'Content-Type': 'application/json',
    },
    ...options,
  })

  const text = await response.text()
  const data = text ? JSON.parse(text) : {}

  if (!response.ok) {
    throw new Error(data?.error || response.statusText || 'Request failed')
  }

  return data as T
}

export async function createJob(repo_url: string) {
  return apiRequest<{ job_id: string }>('/ingest/repo', {
    method: 'POST',
    body: JSON.stringify({ repo_url }),
  })
}

export async function fetchJobs() {
  return apiRequest<{ jobs: { job_id: string; repo_url: string; created_at: string }[] }>('/jobs')
}

export async function fetchSummary(job_id: string) {
  return apiRequest<{ job: { job_id: string; repo_url?: string; status?: string; repo_summary?: RepoSummary | null; owner:string; repo_name:string }  }>(`/result/${job_id}`)
}

export async function queryJob(job_id: string, question: string) {
  return apiRequest<{ answer: string }>('/user-query', {
    method: 'POST',
    body: JSON.stringify({ job_id, question }),
  })
}
