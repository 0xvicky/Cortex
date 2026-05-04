import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { fetchJobs } from '../api'
import type { JobRecord } from '../types'

const repoPattern = /^https:\/\/github\.com\/([^/]+)\/([^/]+)(?:\/)?$/

function formatRepoName(repo_url: string) {
  const match = repoPattern.exec(repo_url)
  return match ? `${match[1]}/${match[2]}` : repo_url
}

function formatDate(iso: string) {
  return new Date(iso).toLocaleString(undefined, {
    dateStyle: 'medium',
    timeStyle: 'short',
  })
}

export default function JobsPage() {
  const [jobs, setJobs] = useState<JobRecord[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const navigate = useNavigate()

  useEffect(() => {
    let active = true

    async function loadJobs() {
      try {
        const data = await fetchJobs()
        if (!active) return
        setJobs(data.jobs || [])
      } catch (err) {
        if (!active) return
        setError(err instanceof Error ? err.message : 'Unable to load jobs')
      } finally {
        if (!active) return
        setLoading(false)
      }
    }

    loadJobs()
    return () => {
      active = false
    }
  }, [])

  return (
    <section className="space-y-10">
      <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
        <h1 className="text-3xl font-semibold leading-tight text-white">Recent analysis jobs</h1>
        <button
          type="button"
          onClick={() => navigate('/')}
          className="inline-flex items-center rounded-[1.5rem] bg-gradient-to-r from-purple-600 to-violet-600 px-6 py-3 text-sm font-semibold text-white shadow-[0_18px_50px_-30px_rgba(124,58,237,0.9)] transition duration-200 hover:shadow-[0_22px_80px_-40px_rgba(124,58,237,0.8)]"
        >
          Analyze a new repo
        </button>
      </div>

      {loading ? (
        <div className="grid gap-5">
          {[1, 2, 3].map((item) => (
            <div key={item} className="h-32 animate-pulse rounded-[2rem] border border-white/10 bg-white/5 p-6" />
          ))}
        </div>
      ) : error ? (
        <div className="rounded-[2rem] border border-rose-500/20 bg-rose-500/10 p-8 text-rose-100">
          <p className="text-lg font-semibold">Unable to load jobs</p>
          <p className="mt-3 text-sm text-rose-200">{error}</p>
        </div>
      ) : jobs.length === 0 ? (
        <div className="rounded-[2rem] border border-white/10 bg-white/5 p-10 text-gray-400">
          <p className="text-2xl font-semibold text-white">No jobs yet.</p>
          <p className="mt-4 max-w-2xl text-base leading-7">Analyze a GitHub repository to generate a job and start asking follow-up questions.</p>
        </div>
      ) : (
        <div className="grid gap-5">
          {jobs.map((job) => (
            <button
              key={job.job_id}
              type="button"
              onClick={() => navigate(`/jobs/${job.job_id}`)}
              className="w-full text-left rounded-[2rem] border border-white/10 bg-white/5 p-8 transition duration-200 hover:border-purple-500/30 hover:shadow-[0_24px_60px_-36px_rgba(124,58,237,0.6)]"
            >
              <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
                <div>
                  <p className="text-sm uppercase tracking-[0.3em] text-gray-400">{formatRepoName(job.repo_url)}</p>
                  <p className="mt-3 max-w-xl text-base leading-7 text-white">{job.repo_url}</p>
                </div>
                <div className="space-y-2 text-right">
                  <p className="text-sm text-gray-400">Job ID <span className="font-semibold text-white">{job.job_id.slice(0, 8)}</span></p>
                  <p className="text-sm text-gray-400">{formatDate(job.created_at)}</p>
                </div>
              </div>
            </button>
          ))}
        </div>
      )}
    </section>
  )
}
