package yarn

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

// RunBeforeBuild attempts to run yarn install if it finds yarn.lock
func (p *Plugin) RunBeforeBuild(ctx context.Context, root string, args []string) error {
	cmd := p.buildCmd(ctx)
	if cmd == nil {
		return nil
	}

	return cmd.Run()
}

// build command will return the command if yarn.lock is found
// otherwise returns nil
func (p *Plugin) buildCmd(ctx context.Context) *exec.Cmd {
	_, err := os.Stat("yarn.lock")
	if os.IsNotExist(err) {
		return nil
	}

	if _, err := os.Stat(".yarnrc.yml"); err == nil {
		return p.execAsYarnPkg(ctx)
	}

	return p.execAsYarnClassic(ctx)
}

func (p *Plugin) execAsYarnClassic(ctx context.Context) *exec.Cmd {
	_, err := os.Stat("yarn.lock")
	if os.IsNotExist(err) {
		return nil
	}

	fmt.Println(">> Running yarn classic install <<<")

	c := exec.CommandContext(ctx, "yarn", "install", "--no-progress")
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout

	return c
}

func (p *Plugin) execAsYarnPkg(ctx context.Context) *exec.Cmd {
	_, err := os.Stat("yarn.lock")
	if os.IsNotExist(err) {
		return nil
	}

	fmt.Println(">> Running yarnpkg install <<<")

	c := exec.CommandContext(ctx, "yarn", "install")
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout

	return c
}
