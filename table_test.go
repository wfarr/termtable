package termtable

import (
	"fmt"
	"testing"
)

func ExampleTermtable() {
	ta := NewTable(nil)
	ta.SetHeader([]string{"LOWERCASE", "UPPERCASE", "NUMBERS"})
	ta.AddRow([]string{"abc", "ABCD", "12345"})
	ta.AddRow([]string{"defg", "EFGHI", "678"})
	ta.AddRow([]string{"hijkl", "JKL", "9000"})
	fmt.Println(ta.Render())

	// Output:
	// +-----------+-----------+---------+
	// | LOWERCASE | UPPERCASE | NUMBERS |
	// +-----------+-----------+---------+
	// | abc       | ABCD      | 12345   |
	// | defg      | EFGHI     | 678     |
	// | hijkl     | JKL       | 9000    |
	// +-----------+-----------+---------+
}

func ExampleLargerTable() {
	t := NewTable(nil)

	t.SetHeader([]string{"Service", "Status", "Details"})
	t.AddRow([]string{"xen3/available_vg_space", "OK", ""})
	t.AddRow([]string{"xen3/bond0", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/bond1", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/cluster_health", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/diskspace", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/drac", "OK", "SSH OK - OpenSSH_3.9p1 (protocol 2.0)\\n"})
	t.AddRow([]string{"xen3/drbd", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/enable_pager", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/firewall", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/gigabit-interfaces", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/hardware", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/heartbeat_links", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/ip_conntrack", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/last_puppet_run", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/mailqueue", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/memory", "OK", ""})
	t.AddRow([]string{"xen3/nrpe", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/ntp", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/ping", "OK", "PING OK - Packet loss = 0%, RTA = 0.32 ms"})
	t.AddRow([]string{"xen3/puppet_env", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/puppet_run_health", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/smtp", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/ssh", "OK", "SSH OK - OpenSSH_5.1p1 Debian-5github8 (protocol 2"})
	t.AddRow([]string{"xen3/syslog-ng", "CRITICAL", "CHECK_NRPE: Error - Could not complete SSL handsha"})
	t.AddRow([]string{"xen3/uptrack", "OK", ""})

	fmt.Println(t.Render())

	// Output:
	// +-------------------------+----------+----------------------------------------------------+
	// | Service                 | Status   | Details                                            |
	// +-------------------------+----------+----------------------------------------------------+
	// | xen3/available_vg_space | OK       |                                                    |
	// | xen3/bond0              | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/bond1              | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/cluster_health     | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/diskspace          | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/drac               | OK       | SSH OK - OpenSSH_3.9p1 (protocol 2.0)\n            |
	// | xen3/drbd               | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/enable_pager       | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/firewall           | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/gigabit-interfaces | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/hardware           | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/heartbeat_links    | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/ip_conntrack       | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/last_puppet_run    | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/mailqueue          | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/memory             | OK       |                                                    |
	// | xen3/nrpe               | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/ntp                | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/ping               | OK       | PING OK - Packet loss = 0%, RTA = 0.32 ms          |
	// | xen3/puppet_env         | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/puppet_run_health  | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/smtp               | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/ssh                | OK       | SSH OK - OpenSSH_5.1p1 Debian-5github8 (protocol 2 |
	// | xen3/syslog-ng          | CRITICAL | CHECK_NRPE: Error - Could not complete SSL handsha |
	// | xen3/uptrack            | OK       |                                                    |
	// +-------------------------+----------+----------------------------------------------------+
}

func BenchmarkRenderBuffer(b *testing.B) {
	b.StopTimer()
	ta := NewTable(nil)
	ta.SetHeader([]string{"LOWERCASE", "UPPERCASE", "NUMBERS"})
	ta.AddRow([]string{"abc", "ABCD", "12345"})
	ta.AddRow([]string{"defg", "EFGHI", "678"})
	ta.AddRow([]string{"hijkl", "JKL", "9000"})
	ta.AddRow([]string{"defg", "EFGHI", "678"})
	ta.AddRow([]string{"hijkl", "JKL", "9000"})
	ta.AddRow([]string{"defg", "EFGHI", "678"})
	ta.AddRow([]string{"hijkl", "JKL", "9000"})
	ta.AddRow([]string{"defg", "EFGHI", "678"})
	ta.AddRow([]string{"hijkl", "JKL", "9000"})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ta.Render()
	}
}
