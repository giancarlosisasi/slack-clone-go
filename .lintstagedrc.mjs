export default {
  '**/*.{js,jsx,ts,tsx}': ['pnpm check'],
  '**/*.go': ['golangci-lint run --fix'],
}