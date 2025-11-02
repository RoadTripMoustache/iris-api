# Contributing to Iris API

Thank you for your interest in contributing to **Iris API** 🎉
We welcome all kinds of contributions — from bug reports and feature suggestions to code improvements and documentation updates.

---

## 🧑‍💻 Development

### Setup

1. Fork the repository and clone your fork locally.
2. Create a new branch from `main` for your changes.

```bash
git checkout -b feature/your-feature-name
```

### Linting

We use [`staticcheck`](https://staticcheck.io/) for Go code linting.
Before committing, make sure your code passes all lint checks:

```bash
staticcheck ./...
```

### Testing

Run the unit tests to ensure your changes do not break existing functionality:

```bash
go test ./...
```

### Mocks & Local Environment

You can run the project locally with mocked Firebase or database services by setting the appropriate environment variables (see the configuration documentation).
This allows development and testing without relying on external dependencies.

---

## 🧾 Code Guidelines

* Write **clean, well-documented code**.
  Each exported function, struct, or method should have a GoDoc comment describing its purpose and behavior.
* Follow existing code style and structure.
* Keep commits focused and meaningful — one logical change per commit.
* Update or add tests when applicable.

---

## 🚀 Submitting Changes

1. Push your branch to your fork.
2. Open a **Pull Request (PR)** against the `main` branch.
3. In your PR description:

    * Explain **why** you’re making the change.
    * Describe **what** it does.
    * Mention any relevant **issues** or **discussions**.

The maintainers will review your PR and may request changes or improvements before merging.

---

## 🧩 Versioning & Releases

* The current version is tracked in **`version.yaml`**.
* All version updates and notable changes must be documented in **`CHANGELOG.md`**.
* Follow [Semantic Versioning](https://semver.org/) when incrementing versions.

---

## 🙌 Acknowledgment

Once your contribution is merged, please **add your name** to the list of contributors in the `README.md` file under the “Contributors” section.
We appreciate your efforts and want to give you proper credit for your work.
