var __index = {"config":{"lang":["en"],"separator":"[\\s\\-]+","pipeline":["stopWordFilter"]},"docs":[{"location":"index.html","title":"identityvalidation Documentation","text":"<p>Hello and welcome to the identityvalidation documentation. While we always want to provide the most comprehensive documentation possible, we thought you may find the below sections a helpful place to get started.</p> <ul> <li>The Getting Started section provides onboarding material</li> <li>The Development header is the best place to get started on developing on top of and with identityvalidation</li> <li>See the Docs section for a comprehensive rundown of identityvalidation capabilities</li> </ul>"},{"location":"index.html#about-identityvalidation","title":"About identityvalidation","text":"<p>identityvalidation provides security operators the tools they need to validate that their security controls and detections are operating as expected. Designed with data-modeling and data-integration needs in mind, identityvalidation can be used on its own as an interactive CLI, orchestrated as part of a broader data pipeline, or leveraged from within the Method Platform.</p> <p>The number of security-relevant Identity Validation resources that identityvalidation can enumerate are constantly growing. For the most up to date listing, please see the documentation here</p> <p>To learn more about identityvalidation, please see the Documentation site for the most detailed information.</p>"},{"location":"index.html#quick-start","title":"Quick Start","text":""},{"location":"index.html#get-identityvalidation","title":"Get identityvalidation","text":"<p>For the full list of available installation options, please see the Installation page. For convenience, here are some of the most commonly used options:</p> <ul> <li><code>docker run methodsecurity/identityvalidation</code></li> <li><code>docker run ghcr.io/method-security/identityvalidation</code></li> <li>Download the latest binary from the Github Releases page</li> <li>Installation documentation</li> </ul>"},{"location":"index.html#authentication","title":"Authentication","text":"<p>TODO</p>"},{"location":"index.html#general-usage","title":"General Usage","text":"<pre><code>identityvalidation &lt;resource&gt; TODO\n</code></pre>"},{"location":"index.html#examples","title":"Examples","text":"<pre><code>identityvalidation TODO\n</code></pre> <pre><code>identityvalidation TODO\n</code></pre>"},{"location":"index.html#contributing","title":"Contributing","text":"<p>Interested in contributing to identityvalidation? Please see our organization wide Contribution page.</p>"},{"location":"index.html#want-more","title":"Want More?","text":"<p>If you're looking for an easy way to tie identityvalidation into your broader cybersecurity workflows, or want to leverage some autonomy to improve your overall security posture, you'll love the broader Method Platform.</p> <p>For more information, visit us here</p>"},{"location":"index.html#community","title":"Community","text":"<p>identityvalidation is a Method Security open source project.</p> <p>Learn more about Method's open source source work by checking out our other projects here or our organization wide documentation here.</p> <p>Have an idea for a Tool to contribute? Open a Discussion here.</p>"},{"location":"community/community.html","title":"Contributing","text":"<p>For more information on how to get involved in the Method community, please see our organization wide documentation:</p> <ul> <li>Discussions</li> <li>Issues</li> <li>Pull Requests</li> </ul>"},{"location":"development/adding.html","title":"Adding a new capability","text":"<p>By design, identityvalidation breaks every unique category into its own top level command. If you are looking to add a brand new capability to the tool, you can take the following steps.</p> <ol> <li>Add a file to <code>cmd/</code> that corresponds to the sub-command name you'd like to add to the <code>identityvalidation</code> CLI</li> <li>You can use <code>cmd/TODO.go</code> as a template</li> <li>Your file needs to be a member function of the <code>identityvalidation</code> struct and should be of the form <code>Init&lt;cmd&gt;Command</code></li> <li>Add a new member to the <code>identityvalidation</code> struct in <code>cmd/root.go</code> that corresponsds to your command name. Remember, the first letter must be capitalized.</li> <li>Call your <code>Init</code> function from <code>main.go</code></li> <li>Add logic to your commands runtime and put it in its own package within <code>internal</code> (e.g., <code>internal/TODO</code>)</li> </ol>"},{"location":"development/principles.html","title":"Project Principles","text":""},{"location":"development/principles.html#pre-run-run-post-run","title":"Pre-run -&gt; Run -&gt; Post-run","text":"<p>In the root command, we set a <code>PersistentPreRunE</code> and <code>PersistentPostRunE</code> function that is responsible for initializing the output format and Signal data (in the pre-run) and then write that data in the proper format (in the post-run).</p> <p>Within the Run command that every command must implement, the output of the collected data needs to be written back to the struct's <code>OutputSignal.Content</code> value in order to be properly written out to the caller.</p>"},{"location":"development/principles.html#cmd-vs-internal","title":"Cmd vs Internal","text":"<p>By design, the functionality within each command should focus around parsing the variety of flags and options that the command may need to control capability, passing off all real logic into internal modules.</p>"},{"location":"development/setup.html","title":"Development Setup","text":""},{"location":"development/setup.html#adding-a-new-capability","title":"Adding a new capability","text":"<p>To add a new capability to identityvalidation, providing new control validation capabilities to security operators everywhere, please see the adding a new capability page.</p>"},{"location":"development/setup.html#setting-up-your-development-environment","title":"Setting up your development environment","text":"<p>If you've just cloned identityvalidation for the first time, welcome to the community! We use Palantir's godel to streamline local development and goreleaser to handle the heavy lifting on the release process.</p> <p>To get started with godel, you can run</p> <pre><code>./godelw verify\n</code></pre> <p>This will run a number of checks for us, including linters, tests, and license checks. We run this command as part of our CI pipeline to ensure the codebase is consistently passing tests.</p>"},{"location":"development/setup.html#building-the-cli","title":"Building the CLI","text":"<p>We can use godel to build our CLI locally by running</p> <pre><code>./godelw build\n</code></pre> <p>You should see output in <code>out/build/identityvalidate/&lt;version&gt;/&lt;os&gt;-&lt;arch&gt;/identityvalidation</code>.</p> <p>If you'd like to clean this output up, you can run</p> <pre><code>./godelw clean\n</code></pre>"},{"location":"development/setup.html#testing-releases-locally","title":"Testing releases locally","text":"<p>We can use goreleaser locally as well to test our builds. As identityvalidation uses cosign to sign our artifacts and Docker containers during our CI pipeline, we'll want to skip this step when running locally.</p> <pre><code>goreleaser release --snapshot --clean --skip sign\n</code></pre> <p>This should output binaries, distributable tarballs/zips, as well as docker images to your local machine's Docker registry.</p>"},{"location":"docs/index.html","title":"Capabilities","text":"<p>identityvalidation checks a number of security controls and capabilities for cybersecurity professionals, enabling them to ensure compliance and protection. Each of the pages below provides a more in-depth look at the IdentityValidation capabilities related to the specified resource.</p>"},{"location":"docs/index.html#top-level-flags","title":"Top Level Flags","text":"<p>identityvalidation has several top level flags that can be used on any subcommand. These include:</p> <pre><code>Flags:\n  -h, --help                 help for methodaws\n  -o, --output string        Output format (signal, json, yaml). Default value is signal (default \"signal\")\n  -f, --output-file string   Path to output file. If blank, will output to STDOUT\n  -q, --quiet                Suppress output\n  -v, --verbose              Verbose output\n</code></pre>"},{"location":"docs/index.html#version-command","title":"Version Command","text":"<p>Run <code>identityvalidation version</code> to get the exact version information for your binary</p>"},{"location":"docs/index.html#output-formats","title":"Output Formats","text":"<p>For more information on the various output formats that are supported by identityvalidation, see the Output Formats page in our organization wide documentation.</p>"},{"location":"getting-started/basic-usage.html","title":"Basic Usage","text":""},{"location":"getting-started/basic-usage.html#binaries","title":"Binaries","text":"<p>Running as a binary means you don't need to do anything additional for identityvalidation to leverage the environment variables you have already exported. You can test that things are working properly by running:</p> <pre><code>identityvalidation TODO\n</code></pre>"},{"location":"getting-started/basic-usage.html#docker","title":"Docker","text":"<pre><code>docker run \\\n TODO\n</code></pre>"},{"location":"getting-started/installation.html","title":"Getting Started","text":"<p>If you are just getting started with identityvalidation, welcome! This guide will walk you through the process of going zero to one with the tool.</p>"},{"location":"getting-started/installation.html#installation","title":"Installation","text":"<p>identityvalidation is provided in several convenient form factors, including statically compiled binary images on a variety of architectures as well as a Docker image for both x86 and ARM machines.</p> <p>If you do not see an architecture that you require, please open a Discussion to propose adding it.</p>"},{"location":"getting-started/installation.html#binaries","title":"Binaries","text":"<p>identityvalidation currently supports statically compiled binaries across the following operating systems and architectures:</p> OS Architecture Linux 386 Linux arm (goarm 7) Linux amd64 Linux arm64 MacOS amd64 MacOS arm64 Windows amd64 <p>The latest binaries can be downloaded directly from Github.</p>"},{"location":"getting-started/installation.html#docker","title":"Docker","text":"<p>Docker images for identityvalidation are hosted in both Github Container Registry as well as on Docker Hub and can be pulled via:</p> <pre><code>docker pull ghcr.io/method-security/identityvalidation\n</code></pre> <pre><code>docker pull methodsecurity/identityvalidation\n</code></pre>"}]}