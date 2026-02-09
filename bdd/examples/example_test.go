package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/zzzep/go-support/bdd"
)

type fakeInterface interface{}
type fakeMock struct {
	mock.Mock
}

func (f *fakeMock) Method() {
	f.Called()
}

func TestPepe(t *testing.T) {
	scenario := bdd.NewScenario()
	f := &fakeMock{}
	scenario.AddMock((*fakeInterface)(nil), f)
	scenario.AddStep(GivenThatIHaveData)
	scenario.AddStep(AndMockWithSuccess)
	scenario.AddStep(WhenVerifyingIfIHaveWhen)
	scenario.AddStep(ThenShouldReturnTrue)
	bdd.NewBDD(t).
		AddScenario(*scenario).
		Run()
}

func TestError(t *testing.T) {
	assert.Panics(t, func() {
		b := bdd.NewBDD(t)
		scenario := bdd.NewScenario()
		scenario.AddStep(func(_ bdd.Scenario, t *testing.T) {
			assert.True(t, false)
		})
		b.AddScenario(*scenario).Run()
	})
}

func GivenThatIHaveData(sc bdd.Scenario, t *testing.T) {
	println("Step 1")
}

func AndMockWithSuccess(sc bdd.Scenario, t *testing.T) {
	sc.Mocks[(*fakeInterface)(nil)].(*fakeMock).On("Method").Return().Once()
	sc.Mocks[(*fakeInterface)(nil)].(*fakeMock).Method()
	println("Step Mock")
}

func WhenVerifyingIfIHaveWhen(sc bdd.Scenario, t *testing.T) {
	println("Step 2")
}

func ThenShouldReturnTrue(sc bdd.Scenario, t *testing.T) {
	println("Step 3")
	assert.True(t, true)
}
