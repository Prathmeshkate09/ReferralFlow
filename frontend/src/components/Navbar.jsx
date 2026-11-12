import React from 'react'

export default function Navbar({ unlocked, score }) {
  return (
    <header className="sticky top-0 z-30 bg-white/95 backdrop-blur border-b border-slate-200">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 h-16 flex items-center justify-between">
        <div className="flex items-center gap-3">
          <div className="h-9 w-9 rounded-xl bg-indigo-600 flex items-center justify-center text-white font-bold text-lg shadow-sm">
            R
          </div>
          <div>
            <span className="font-bold text-lg text-slate-900 tracking-tight">Referral<span className="text-indigo-600">Flow</span></span>
            <span className="hidden sm:inline-block ml-2 text-xs bg-indigo-50 text-indigo-700 font-medium px-2 py-0.5 rounded-full border border-indigo-200">
              Verified Jobs & Fast Referrals
            </span>
          </div>
        </div>

        <div className="flex items-center gap-4">
          <div className="flex items-center gap-2 bg-slate-50 border border-slate-200 rounded-lg px-3 py-1.5 text-xs">
            <span className="text-slate-500">Profile Completion:</span>
            <span className="font-semibold text-slate-900">{score}%</span>
            {unlocked ? (
              <span className="bg-emerald-100 text-emerald-800 text-[10px] font-bold px-1.5 py-0.5 rounded">
                Unlocked
              </span>
            ) : (
              <span className="bg-amber-100 text-amber-800 text-[10px] font-bold px-1.5 py-0.5 rounded">
                Locked (&lt;80%)
              </span>
            )}
          </div>
          <a
            href="https://github.com/Prathmeshkate09/ReferralFlow"
            target="_blank"
            rel="noreferrer"
            className="text-xs font-medium text-slate-600 hover:text-indigo-600 transition"
          >
            GitHub
          </a>
        </div>
      </div>
    </header>
  )
}
