import * as React from 'react';
import { NextPage } from 'next';

const IndexPage: NextPage = () => {
  return (
    <>
      <header className="h-16">
        <div className="mx-auto w-full h-full max-w-screen-md px-8 flex justify-between items-stretch">
          <div className="flex items-center">
            <span className="font-bold text-3xl">
              Letterbookd
            </span>
          </div>
          <div className="flex items-center gap-8">
            <form>
              <button
                type="button"
                className="h-12 rounded-full overflow-hidden relative px-4 uppercase font-bold"
              >
                <span className="absolute top-0 left-0 w-full h-full rounded-[inherit] border-2" />
                <span className="relative flex items-center gap-2">
                  <span className="rounded-full w-8 h-8 flex items-center justify-center -ml-2 relative overflow-hidden">
                    <span className="absolute top-0 left-0 w-full h-full bg-current opacity-25" />
                    <span className="relative">
                      I
                    </span>
                  </span>
                  <span>
                    Sign Up
                  </span>
                </span>
              </button>
            </form>
            <form>
              <div className="h-12 rounded-full overflow-hidden relative">
                <span className="absolute top-0 left-0 border-2 rounded-[inherit] w-full h-full pointer-events-none" />
                <span className="w-12 h-full absolute top-0 right-0 flex items-center justify-center">
                  <svg
                    className="relative -top-0.5 -left-1"
                    viewBox="0 0 24 24"
                    width="24"
                    height="24"
                    fill="none"
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="2"
                  >
                    <circle
                      cx="11"
                      cy="11"
                      r="8"
                    />
                    <line
                      x1="21"
                      x2="16.65"
                      y1="21"
                      y2="16.65"
                    />
                  </svg>
                </span>
                <input
                  type="search"
                  className="h-full bg-background pl-4 pr-12"
                />
              </div>
            </form>
          </div>
        </div>
      </header>
      <main>
        <div className="py-32">
          <div className="mx-auto w-full h-full max-w-screen-md px-8">
            <p className="text-5xl font-bold text-center">
              <strong className="font-bold">
                Books: Read, Review, Recommend
              </strong>
            </p>
            <div className="mt-16 text-center">
              <a
                className="inline-flex px-8 rounded h-12 relative items-center justify-center text-center font-bold"
                href="/signup"
              >
                <span className="absolute top-0 left-0 w-full h-full rounded-[inherit] border-2" />
                <span>
                  Get started&mdash;it's free!
                </span>
              </a>
            </div>
          </div>
        </div>
        <div className="py-32">
          <div className="mx-auto w-full h-full max-w-screen-md px-8">
            <section>
              <h1>
                Featured
              </h1>
              <ul className="grid grid-cols-4 gap-2 h-48">
                <li className="h-full">
                  <a
                    href="#"
                    className="block h-full"
                  >
                    <img
                      src="http://placehold.it/1"
                      alt="foo"
                      className="w-full h-full object-center object-cover rounded"
                    />
                  </a>
                </li>
                <li className="h-full">
                  <a
                    href="#"
                    className="block h-full"
                  >
                    <img
                      src="http://placehold.it/1"
                      alt="foo"
                      className="w-full h-full object-center object-cover rounded"
                    />
                  </a>
                </li>
                <li className="h-full">
                  <a
                    href="#"
                    className="block h-full"
                  >
                    <img
                      src="http://placehold.it/1"
                      alt="foo"
                      className="w-full h-full object-center object-cover rounded"
                    />
                  </a>
                </li>
                <li className="h-full">
                  <a
                    href="#"
                    className="block h-full"
                  >
                    <img
                      src="http://placehold.it/1"
                      alt="foo"
                      className="w-full h-full object-center object-cover rounded"
                    />
                  </a>
                </li>
              </ul>
            </section>
          </div>
        </div>
      </main>
      <footer>
        <div className="mx-auto w-full h-full max-w-screen-md px-8 py-16">
          <p className="text-center">
            Copyright &copy; 2024.
          </p>
        </div>
      </footer>
    </>
);
}

export default IndexPage;
