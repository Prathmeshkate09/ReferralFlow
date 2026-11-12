import React from 'react'

export default function StatsBanner() {
  return (
    <div className="bg-gradient-to-r from-indigo-900 via-slate-900 to-indigo-950 text-white rounded-2xl p-6 sm:p-8 shadow-md mb-8">
      <div className="max-w-3xl">
        <div className="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-indigo-500/20 text-indigo-300 text-xs font-semibold mb-3 border border-indigo-500/30">
          <span>⚡ 24-Hour Referral Delivery</span>
        </div>
        <h1 className="text-2xl sm:text-3xl font-bold tracking-tight mb-2">
          Find Verified Jobs. Get Directly Referred.
        </h1>
        <p className="text-slate-300 text-sm sm:text-base leading-relaxed">
          Skip candidate black holes. ReferralFlow matches verified job seekers with active referrers using automated Go-powered state routing and PostgreSQL index-optimized discovery.
        </p>
      </div>

      <div className="grid grid-cols-2 sm:grid-cols-4 gap-4 mt-6 pt-6 border-t border-slate-700/60">
        <div>
          <div className="text-xl sm:text-2xl font-bold text-indigo-400">100%</div>
          <div className="text-xs text-slate-400">Verified Opportunities</div>
        </div>
        <div>
          <div className="text-xl sm:text-2xl font-bold text-emerald-400">&lt; 24h</div>
          <div className="text-xs text-slate-400">Referral Turnaround</div>
        </div>
        <div>
          <div className="text-xl sm:text-2xl font-bold text-amber-400">80%</div>
          <div className="text-xs text-slate-400">Profile Unlock Gate</div>
        </div>
        <div>
          <div className="text-xl sm:text-2xl font-bold text-sky-400">Go + React</div>
          <div className="text-xs text-slate-400">Modern Architecture</div>
        </div>
      </div>
    </div>
  )
}
