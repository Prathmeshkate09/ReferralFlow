import React, { useState } from 'react'

export default function ReferralModal({ job, onClose, onSubmitReferral }) {
  const [pitch, setPitch] = useState(
    'I have built production services with Go, React, and PostgreSQL, and I would love a referral for this opportunity.'
  )
  const [submitting, setSubmitting] = useState(false)

  if (!job) return null

  const handleSubmit = async (e) => {
    e.preventDefault()
    setSubmitting(true)
    await onSubmitReferral(job.id, pitch)
    setSubmitting(false)
  }

  return (
    <div className="fixed inset-0 z-50 bg-slate-900/60 backdrop-blur-sm flex items-center justify-center p-4">
      <div className="bg-white rounded-2xl max-w-lg w-full p-6 shadow-xl border border-slate-100">
        <div className="flex items-start justify-between mb-4">
          <div>
            <span className="text-xs font-semibold text-indigo-600 uppercase tracking-wider">
              Request Referral in 24h
            </span>
            <h2 className="text-lg font-bold text-slate-900 mt-1">{job.title}</h2>
            <p className="text-xs text-slate-500">{job.company} • {job.workplace}</p>
          </div>
          <button
            onClick={onClose}
            className="text-slate-400 hover:text-slate-600 text-xl font-bold p-1"
          >
            ×
          </button>
        </div>

        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="block text-xs font-semibold text-slate-700 mb-1">
              Your Pitch to the Referrer:
            </label>
            <textarea
              rows={4}
              value={pitch}
              onChange={(e) => setPitch(e.target.value)}
              className="w-full text-xs p-3 border border-slate-200 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:outline-none"
              placeholder="Highlight your relevant projects, tech stack, and why you're a great fit..."
              required
            />
          </div>

          <div className="bg-indigo-50 border border-indigo-100 rounded-lg p-3 text-xs text-indigo-800">
            <strong>How it works:</strong> Your verified profile details (resume, GitHub, technical skills) and pitch will be dispatched to an active employee referrer. You will receive referral status confirmation.
          </div>

          <div className="flex items-center justify-end gap-3 pt-3">
            <button
              type="button"
              onClick={onClose}
              className="text-xs font-medium text-slate-600 hover:text-slate-800 px-3 py-2 rounded-lg"
            >
              Cancel
            </button>
            <button
              type="submit"
              disabled={submitting}
              className="bg-indigo-600 hover:bg-indigo-700 text-white text-xs font-semibold px-4 py-2 rounded-lg transition disabled:opacity-50 shadow-sm"
            >
              {submitting ? 'Submitting Request...' : 'Confirm Referral Request'}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}
