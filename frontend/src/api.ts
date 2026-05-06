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

export async function createJob(repoUrl: string, userId:string) {
  return apiRequest<{ job_id: string }>('/ingest/repo', {
    method: 'POST',
    body: JSON.stringify({ repoUrl,userId }),
  })
}

export async function fetchJobs(userId:string) {
  return apiRequest<{ res: { job_id: string}[] }>(`/jobs/${userId}`)
}

export async function fetchSummary(job_id: string,user_id:string) {
    console.log(user_id, job_id)
  return apiRequest<{ job: { job_id: string; repo_url?: string; status?: string; repo_summary?: RepoSummary | null;  }  }>(`/result/${user_id}/${job_id}`)
}

export async function queryJob(jobId: string, userQuery: string) {
  return apiRequest<{ answer: string }>('/user-query', {
    method: 'POST',
    body: JSON.stringify({ jobId, userQuery }),
  })
}
